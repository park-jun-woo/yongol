//ff:func feature=validate-contract type=util control=sequence
//ff:what canonicalFieldKey — DDL 컬럼 / Go 필드 이름을 소문자·무언더스코어 키로 정규화

package contract

import "strings"

// canonicalFieldKey lower-cases the name and strips underscores so
// `OrgID`, `orgId`, `org_id`, `ORG_ID` all map to the same key
// `orgid`. Callers pass both DDL column names (snake_case) and Go
// struct field names (PascalCase with initialisms) — the canonical
// form absorbs both styles and tolerates sqlc / oapi-codegen
// initialism rewrites.
func canonicalFieldKey(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), "_", "")
}
