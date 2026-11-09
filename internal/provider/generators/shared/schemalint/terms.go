package schemalint

import (
	"fmt"
	"sort"
	"strings"
)

// term is an entry of the fixed spelling table. Keys match property names
// case-insensitively; new fields must use the canonical spelling.
type term struct {
	key       string
	canonical string
}

// terms is the fixed spelling table, longest key first so that overlapping
// terms resolve to the complete word or phrase (Ipv6 before Ip, IpAddress
// before Ip). Plurals are derived by appending a lower-case "s" (IDs -> Ids).
var terms = func() []term {
	t := []term{
		{key: "ip", canonical: "Ip"},
		{key: "ipv4", canonical: "Ipv4"},
		{key: "ipv6", canonical: "Ipv6"},
		{key: "vpc", canonical: "Vpc"},
		{key: "cidr", canonical: "Cidr"},
		{key: "eip", canonical: "Eip"},
		{key: "id", canonical: "Id"},
		{key: "name", canonical: "Name"},
		{key: "description", canonical: "Description"},
		{key: "desc", canonical: "Description"},
		{key: "status", canonical: "Status"},
		{key: "tags", canonical: "Tags"},
		{key: "ipaddress", canonical: "IpAddress"},
		{key: "projectname", canonical: "ProjectName"},
		{key: "securitygroup", canonical: "SecurityGroup"},
	}
	sort.SliceStable(t, func(i, j int) bool { return len(t[i].key) > len(t[j].key) })
	return t
}()

// reversedPhrases maps adjacent words written in reverse order to the fixed
// word order of the phrase, e.g. AddressIp -> IpAddress.
var reversedPhrases = map[[2]string][2]string{
	{"address", "ip"}:     {"Ip", "Address"},
	{"name", "project"}:   {"Project", "Name"},
	{"group", "security"}: {"Security", "Group"},
}

// token is a word of a property name.
type token struct {
	text      string // spelling in the property name
	canonical string // fixed spelling for table terms, otherwise text
	term      bool   // whether the token is a table term
}

// analysis is the result of checking a property name against the fixed
// spelling table.
type analysis struct {
	canonical     string   // suggested property name
	spellingFixes []string // e.g. `"IDs" -> "Ids"`
	orderFixes    []string // e.g. `"AddressIp" -> "IpAddress"`
}

// analyzeName rewrites table terms in a property name to their fixed spelling
// and word order.
func analyzeName(name string) analysis {
	var result analysis

	tokens := tokenize(name)
	words := make([]string, len(tokens))
	for i, t := range tokens {
		words[i] = t.canonical
		if t.term && t.text != t.canonical {
			result.spellingFixes = append(result.spellingFixes, fmt.Sprintf("%q -> %q", t.text, t.canonical))
		}
	}

	for i := 0; i+1 < len(words); i++ {
		fixed, ok := reversedPhrases[[2]string{strings.ToLower(words[i]), strings.ToLower(words[i+1])}]
		if !ok {
			continue
		}
		result.orderFixes = append(result.orderFixes, fmt.Sprintf("%q -> %q", words[i]+words[i+1], fixed[0]+fixed[1]))
		words[i], words[i+1] = fixed[0], fixed[1]
		i++
	}

	var sb strings.Builder
	for _, w := range words {
		if !isAlphanumeric(w) {
			continue // Drop separators such as '_' and '-'.
		}
		sb.WriteString(strings.ToUpper(w[:1]) + w[1:])
	}
	result.canonical = sb.String()

	return result
}

// tokenize splits a property name into words. Table terms are recognized
// first; an upper-case run that cannot be split entirely into table terms is
// kept as one out-of-table acronym, so VIP, IDC or UUID are never read as
// V+IP, ID+C or UU+ID.
func tokenize(name string) []token {
	var tokens []token

	for i := 0; i < len(name); {
		if block, end, ok := matchTermBlock(name, i); ok {
			tokens = append(tokens, block...)
			i = end
			continue
		}

		j := i + 1
		switch c := name[i]; {
		case isDigit(c):
			for j < len(name) && isDigit(name[j]) {
				j++
			}
		case isLower(c):
			for j < len(name) && isLower(name[j]) {
				j++
			}
		case isUpper(c):
			for j < len(name) && isUpper(name[j]) {
				j++
			}
			if j-i == 1 {
				// Capitalized word, e.g. "Instance".
				for j < len(name) && isLower(name[j]) {
					j++
				}
			} else if j < len(name) && isLower(name[j]) {
				// Acronym followed by a capitalized word, e.g. "DB" in "DBName".
				j--
			}
		}

		tokens = append(tokens, token{text: name[i:j], canonical: name[i:j]})
		i = j
	}

	return tokens
}

// matchTermBlock matches consecutive table terms starting at i. The block must
// not end inside an upper-case run unless the rest of the run is the first
// letter of a capitalized word (IPAddress, IDCard).
func matchTermBlock(name string, i int) ([]token, int, bool) {
	if i > 0 && !isUpper(name[i]) {
		return nil, 0, false
	}

	runEnd := i
	for runEnd < len(name) && isUpper(name[runEnd]) {
		runEnd++
	}

	var block []token
	pos := i
	for pos < len(name) && (pos == i || isUpper(name[pos])) {
		t, end, ok := matchTerm(name, pos)
		if !ok {
			break
		}
		block = append(block, t)
		pos = end
		if pos >= runEnd {
			break
		}
	}

	if len(block) == 0 {
		return nil, 0, false
	}

	if pos < runEnd && (runEnd-pos != 1 || runEnd == len(name) || !isLower(name[runEnd])) {
		return nil, 0, false
	}

	return block, pos, true
}

// matchTerm matches a single table term (or its plural) at pos. The term must
// end at a word boundary: the end of the name, an upper-case letter or a digit.
func matchTerm(name string, pos int) (token, int, bool) {
	for _, t := range terms {
		for _, plural := range []bool{true, false} {
			key, canonical := t.key, t.canonical
			if plural {
				key, canonical = key+"s", canonical+"s"
			}

			end := pos + len(key)
			if end > len(name) || !strings.EqualFold(name[pos:end], key) {
				continue
			}
			if plural && name[end-1] != 's' {
				continue
			}
			if end < len(name) && !isUpper(name[end]) && !isDigit(name[end]) {
				continue
			}

			return token{text: name[pos:end], canonical: canonical, term: true}, end, true
		}
	}

	return token{}, 0, false
}

// isPascalCase reports whether name starts with an upper-case letter and
// contains only ASCII letters and digits.
func isPascalCase(name string) bool {
	return name != "" && isUpper(name[0]) && isAlphanumeric(name)
}

func isAlphanumeric(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !isUpper(s[i]) && !isLower(s[i]) && !isDigit(s[i]) {
			return false
		}
	}
	return true
}

func isUpper(c byte) bool { return c >= 'A' && c <= 'Z' }
func isLower(c byte) bool { return c >= 'a' && c <= 'z' }
func isDigit(c byte) bool { return c >= '0' && c <= '9' }
