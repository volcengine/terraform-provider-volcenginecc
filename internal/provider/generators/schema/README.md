# Volcengine Cloud Control Resource Schema Downloader

## Schema naming validation

Before any code is generated, every configured schema is checked against the
schema naming rules (`internal/provider/generators/shared/schemalint`). A single
violation fails the whole run and no file is written. The resource and data
source generators run the same check.

The rules cover root and nested properties:

| Rule | Requirement |
| --- | --- |
| `pascal-case` | Field names are PascalCase. |
| `term-spelling` | Table terms use their fixed spelling: `Ip`, `Ipv4`, `Ipv6`, `Vpc`, `Cidr`, `Eip`, `Id`, `Ids`, `Name`, `Description` (not `Desc`), `Status`, `Tags`, `IpAddress`, `ProjectName`, `SecurityGroup`. Overlapping terms match the longest word or phrase first, e.g. `IPv6Address` -> `Ipv6Address`, `SecurityGroupIDs` -> `SecurityGroupIds`. |
| `term-order` | Phrases keep their fixed word order, e.g. `AddressIp` -> `IpAddress`. |
| `term-consistency` | An out-of-table term (e.g. `DB`, `Http`) is spelled one way at every level of a resource and reuses the spelling of published fields. |
| `duplicate-field` | Spelling variants of one field are not declared side by side, e.g. `IPAddress` and `IpAddress`. |
| `json-pointer` | `required`, identifiers, property lists, `filterProperties` and `tagging.tagProperty` reference existing field names exactly; array elements are addressed with `*`. |
| `tag-property` | Taggable resources declare tags as the root property `Tags`. |

Descriptions are not checked, so OpenAPI and SDK parameter names can be quoted
with their original spelling.

### Naming baseline

Published fields that predate the rules keep their full names. They are listed
per Cloud Control type in
`internal/provider/generators/shared/schemalint/naming_baseline.json` and are
exempt from the checks. Entries are only ever removed: new fields must follow
the rules instead of being added to the baseline. `go test
./internal/provider/generators/shared/schemalint/` fails on entries that no
longer match a violation so they can be deleted.
