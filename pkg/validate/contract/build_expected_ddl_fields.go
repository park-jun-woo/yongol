//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what buildExpectedDDLFields — DDLTables 컬럼명을 canonical(소문자·무언더스코어) 키로 집합화

package contract

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildExpectedDDLFields returns the set of DDL column names reduced to
// a canonical form — lower-case, underscores removed — so the match
// against Go struct field names tolerates initialism variants emitted
// by sqlc / oapi-codegen (e.g. `org_id` → `OrgID`, `id` → `ID`).
//
// PRV-02 normalizes each observed `recv.Field` selector to the same
// canonical form before looking it up, removing a whole class of false
// positives rooted in capitalization style.
func buildExpectedDDLFields(fs *yongol.Fullstack) map[string]bool {
	out := map[string]bool{}
	for _, t := range fs.DDLTables {
		for col := range t.Columns {
			out[canonicalFieldKey(col)] = true
		}
	}
	return out
}
