// Package schemalint validates Cloud Control resource schema property names
// against the schema naming rules before Terraform code is generated:
//
//   - Property names are PascalCase.
//   - Table terms use their fixed spelling and word order (Ip, Ipv6, Vpc, Cidr,
//     Eip, Id, Ids, Name, Description, Status, Tags, IpAddress, ProjectName,
//     SecurityGroup).
//   - Out-of-table terms are spelled the same way at every level of a resource.
//   - Spelling variants of one field are not declared side by side.
//   - JSON Pointers (required, identifiers, property lists, filters, tagging)
//     reference existing schema field names exactly.
//   - Taggable resources expose tags as the root property "Tags".
//
// Published fields that predate these rules keep their names; they are listed
// in the naming baseline and exempt from the checks.
package schemalint

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/volcengine/terraform-provider-volcenginecc/internal/ccschema"
)

// Rule identifies a schema naming rule.
type Rule string

const (
	RulePascalCase      Rule = "pascal-case"
	RuleTermSpelling    Rule = "term-spelling"
	RuleTermOrder       Rule = "term-order"
	RuleTermConsistency Rule = "term-consistency"
	RuleDuplicateField  Rule = "duplicate-field"
	RuleJSONPointer     Rule = "json-pointer"
	RuleTagProperty     Rule = "tag-property"
)

const (
	propertiesPrefix = "/properties"
	tagsPointer      = propertiesPrefix + "/Tags"
	taggingPath      = "/tagging"
)

var arrayIndex = regexp.MustCompile(`^[0-9]+$`)

// Violation is a property name or JSON Pointer that breaks a naming rule.
type Violation struct {
	TypeName string
	// Path is the JSON Pointer of the offending field, the offending JSON
	// Pointer value itself, or "/tagging". It is the key used by the baseline.
	Path    string
	Rule    Rule
	Message string
}

func (v Violation) String() string {
	return fmt.Sprintf("%s: %s [%s]", v.Path, v.Message, v.Rule)
}

// Error reports all violations of one resource type.
type Error struct {
	TypeName   string
	Violations []Violation
}

func (e *Error) Error() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s violates %d schema naming rule(s):", e.TypeName, len(e.Violations))
	for _, v := range e.Violations {
		sb.WriteString("\n  ")
		sb.WriteString(v.String())
	}
	return sb.String()
}

// Check validates an expanded resource schema and returns an *Error if any
// field outside the baseline violates a naming rule.
func Check(resource *ccschema.Resource, baseline Baseline) error {
	violations := Validate(resource, baseline)
	if len(violations) == 0 {
		return nil
	}

	return &Error{TypeName: violations[0].TypeName, Violations: violations}
}

// Validate validates an expanded resource schema (see ccschema.Resource.Expand)
// and returns the violations of fields outside the baseline, sorted by path.
func Validate(resource *ccschema.Resource, baseline Baseline) []Violation {
	if resource == nil {
		return nil
	}

	v := &validator{exempt: make(map[string]bool), seen: make(map[Violation]bool)}
	if resource.TypeName != nil {
		v.typeName = *resource.TypeName
	}
	for _, path := range baseline[v.typeName] {
		v.exempt[path] = true
	}

	fields := collectFields(resource.Properties, propertiesPrefix, nil)
	v.checkNames(fields)
	v.checkTermConsistency(fields)
	v.checkDuplicateFields(fields)
	v.checkPointers(resource)
	v.checkRequiredLists(resource, fields)
	v.checkTagging(resource)

	sort.SliceStable(v.violations, func(i, j int) bool {
		if v.violations[i].Path != v.violations[j].Path {
			return v.violations[i].Path < v.violations[j].Path
		}
		return v.violations[i].Rule < v.violations[j].Rule
	})

	return v.violations
}

// field is a property declared in the schema.
type field struct {
	name     string
	path     string // e.g. /properties/Rules/*/Name
	parent   string // e.g. /properties/Rules/*
	property *ccschema.Property
}

// collectFields walks the expanded properties in stable name order. Array items
// are addressed with "*".
func collectFields(properties map[string]*ccschema.Property, parent string, fields []field) []field {
	for _, name := range sortedNames(properties) {
		property := properties[name]
		path := parent + "/" + name
		fields = append(fields, field{name: name, path: path, parent: parent, property: property})
		if property == nil {
			continue
		}

		fields = collectFields(property.Properties, path, fields)
		itemsPath := path
		for items := property.Items; items != nil; items = items.Items {
			itemsPath += "/*"
			fields = collectFields(items.Properties, itemsPath, fields)
		}
	}

	return fields
}

type validator struct {
	typeName   string
	exempt     map[string]bool
	violations []Violation
	seen       map[Violation]bool
}

func (v *validator) add(path string, rule Rule, format string, a ...any) {
	if v.exempt[path] {
		return
	}

	violation := Violation{
		TypeName: v.typeName,
		Path:     path,
		Rule:     rule,
		Message:  fmt.Sprintf(format, a...),
	}
	if v.seen[violation] {
		return
	}
	v.seen[violation] = true
	v.violations = append(v.violations, violation)
}

// checkNames enforces PascalCase and the fixed spelling and word order of
// table terms.
func (v *validator) checkNames(fields []field) {
	for _, f := range fields {
		a := analyzeName(f.name)

		if !isPascalCase(f.name) {
			v.add(f.path, RulePascalCase, "field name %q must be PascalCase; suggested name %q", f.name, a.canonical)
		}
		if len(a.spellingFixes) > 0 {
			v.add(f.path, RuleTermSpelling, "use the fixed term spelling %s; suggested name %q", strings.Join(a.spellingFixes, ", "), a.canonical)
		}
		if len(a.orderFixes) > 0 {
			v.add(f.path, RuleTermOrder, "use the fixed word order %s; suggested name %q", strings.Join(a.orderFixes, ", "), a.canonical)
		}
	}
}

// checkTermConsistency requires an out-of-table term to be spelled one way at
// every level of the resource. Spellings used by baseline fields are published
// and may be reused; any other spelling must agree with the rest of the
// resource.
func (v *validator) checkTermConsistency(fields []field) {
	type usage struct {
		spelling string
		path     string
	}

	published := make(map[string]map[string]bool)   // term -> spellings of baseline fields
	firstPath := make(map[string]map[string]string) // term -> spelling -> first field using it
	uses := make(map[string][]usage)

	for _, f := range fields {
		for _, t := range tokenize(f.name) {
			if t.term || len(t.text) < 2 || !isUpper(t.text[0]) {
				// Table terms are covered by checkNames, lower-case words by PascalCase.
				continue
			}

			key := strings.ToLower(t.text)
			if firstPath[key] == nil {
				firstPath[key] = make(map[string]string)
			}
			if _, ok := firstPath[key][t.text]; !ok {
				firstPath[key][t.text] = f.path
			}

			if v.exempt[f.path] {
				if published[key] == nil {
					published[key] = make(map[string]bool)
				}
				published[key][t.text] = true
				continue
			}
			uses[key] = append(uses[key], usage{spelling: t.text, path: f.path})
		}
	}

	for _, key := range sortedNames(uses) {
		if len(firstPath[key]) < 2 {
			continue
		}

		for _, u := range uses[key] {
			if published[key][u.spelling] {
				continue
			}

			var others []string
			for _, spelling := range sortedNames(firstPath[key]) {
				if spelling != u.spelling {
					others = append(others, fmt.Sprintf("%q (%s)", spelling, firstPath[key][spelling]))
				}
			}

			if len(published[key]) > 0 {
				v.add(u.path, RuleTermConsistency, "term %q must reuse the published spelling %s", u.spelling, strings.Join(others, ", "))
			} else {
				v.add(u.path, RuleTermConsistency, "term %q is also spelled %s in this resource; use one spelling at every level", u.spelling, strings.Join(others, ", "))
			}
		}
	}
}

// checkDuplicateFields rejects sibling fields that are spelling variants of
// the same schema field, e.g. IPAddress and IpAddress.
func (v *validator) checkDuplicateFields(fields []field) {
	type group struct {
		parent, canonical string
	}

	groups := make(map[group][]field)
	var order []group
	for _, f := range fields {
		g := group{parent: f.parent, canonical: analyzeName(f.name).canonical}
		if _, ok := groups[g]; !ok {
			order = append(order, g)
		}
		groups[g] = append(groups[g], f)
	}

	for _, g := range order {
		variants := groups[g]
		if len(variants) < 2 {
			continue
		}

		for _, f := range variants {
			var others []string
			for _, other := range variants {
				if other.name != f.name {
					others = append(others, fmt.Sprintf("%q", other.name))
				}
			}
			v.add(f.path, RuleDuplicateField, "%q and %s are spelling variants of %q; map them to one schema field", f.name, strings.Join(others, ", "), g.canonical)
		}
	}
}

// checkPointers requires every JSON Pointer to resolve to a declared field with
// the exact schema field name.
func (v *validator) checkPointers(r *ccschema.Resource) {
	type pointerList struct {
		keyword  string
		pointers ccschema.PropertyJsonPointers
	}

	lists := []pointerList{
		{"primaryIdentifier", r.PrimaryIdentifier},
		{"readOnlyProperties", r.ReadOnlyProperties},
		{"createOnlyProperties", r.CreateOnlyProperties},
		{"conditionalCreateOnlyProperties", r.ConditionalCreateOnlyProperties},
		{"writeOnlyProperties", r.WriteOnlyProperties},
		{"deprecatedProperties", r.DeprecatedProperties},
		{"nonPublicProperties", r.NonPublicProperties},
	}
	for i, identifier := range r.AdditionalIdentifiers {
		lists = append(lists, pointerList{fmt.Sprintf("additionalIdentifiers[%d]", i), identifier})
	}
	if r.FilterProperties != nil {
		lists = append(lists,
			pointerList{"filterProperties.filterableProperties", r.FilterProperties.FilterableProperties},
			pointerList{"filterProperties.required", r.FilterProperties.Required},
		)
	}
	if r.Tagging != nil && r.Tagging.TagProperty != nil {
		lists = append(lists, pointerList{"tagging.tagProperty", ccschema.PropertyJsonPointers{*r.Tagging.TagProperty}})
	}

	for _, list := range lists {
		for _, pointer := range list.pointers {
			if reason := resolvePointer(r.Properties, string(pointer)); reason != "" {
				v.add(string(pointer), RuleJSONPointer, "%s pointer %q %s", list.keyword, pointer, reason)
			}
		}
	}
}

// checkRequiredLists checks the root required list and the required lists of
// nested objects and array items.
func (v *validator) checkRequiredLists(r *ccschema.Resource, fields []field) {
	v.checkRequired(propertiesPrefix, r.Required, r.Properties)

	for _, f := range fields {
		if f.property == nil {
			continue
		}
		v.checkRequired(f.path, f.property.Required, f.property.Properties)
		itemsPath := f.path
		for items := f.property.Items; items != nil; items = items.Items {
			itemsPath += "/*"
			v.checkRequired(itemsPath, items.Required, items.Properties)
		}
	}
}

// checkRequired requires each name in a required list to be a sibling field.
func (v *validator) checkRequired(parent string, required []string, properties map[string]*ccschema.Property) {
	for _, name := range required {
		if _, ok := properties[name]; ok {
			continue
		}
		if match := matchFold(properties, name); match != "" {
			v.add(parent+"/"+name, RuleJSONPointer, "required field %q does not match the field name %q", name, match)
		} else {
			v.add(parent+"/"+name, RuleJSONPointer, "required field %q is not declared", name)
		}
	}
}

// resolvePointer returns why a property JSON Pointer does not resolve, or "".
func resolvePointer(properties map[string]*ccschema.Property, pointer string) string {
	if !strings.HasPrefix(pointer, propertiesPrefix+"/") {
		return fmt.Sprintf("must start with %q", propertiesPrefix+"/")
	}

	property := &ccschema.Property{Properties: properties}
	for _, segment := range strings.Split(strings.TrimPrefix(pointer, propertiesPrefix+"/"), "/") {
		if segment == "*" {
			if property.Items == nil {
				return `uses "*" on a field that is not an array`
			}
			property = property.Items
			continue
		}
		if property.Items != nil {
			if arrayIndex.MatchString(segment) {
				return fmt.Sprintf(`addresses array elements by index %q; use "*"`, segment)
			}
			return fmt.Sprintf(`must address the elements of an array with "*" before %q`, segment)
		}

		next, ok := property.Properties[segment]
		if !ok || next == nil {
			if match := matchFold(property.Properties, segment); match != "" {
				return fmt.Sprintf("segment %q does not match the field name %q", segment, match)
			}
			return fmt.Sprintf("segment %q is not a declared field", segment)
		}
		property = next
	}

	return ""
}

// checkTagging requires taggable resources to expose tags as the root
// property "Tags".
func (v *validator) checkTagging(r *ccschema.Resource) {
	if r.Tagging == nil || r.Tagging.Taggable == nil || !*r.Tagging.Taggable {
		return
	}

	if p := r.Tagging.TagProperty; p != nil && string(*p) != tagsPointer {
		v.add(taggingPath, RuleTagProperty, "tagProperty must be %q, got %q", tagsPointer, *p)
	}
	if _, ok := r.Properties["Tags"]; !ok {
		v.add(taggingPath, RuleTagProperty, `taggable resources must declare tags as the root property "Tags"`)
	}
}

// matchFold returns the declared name that equals name case-insensitively.
func matchFold(properties map[string]*ccschema.Property, name string) string {
	for _, declared := range sortedNames(properties) {
		if strings.EqualFold(declared, name) {
			return declared
		}
	}
	return ""
}

func sortedNames[V any](m map[string]V) []string {
	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
