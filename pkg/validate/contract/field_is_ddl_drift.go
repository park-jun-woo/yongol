//ff:func feature=validate-contract type=util control=sequence
//ff:what fieldIsDDLDrift — "recv.Field" 문자열에서 Field 부분만 떼어 DDL 컬럼 세트와 대조

package contract

import "strings"

// fieldDenylist holds Field names that the symbol extractor picks up
// as selector expressions but are structural accessors on the
// generated Server / Queries / transaction plumbing rather than
// DDL columns. Including them in PRV-02 would produce false
// positives on every preserved file that touches `server.DB` or
// `server.Queries`, so we deliberately ignore them.
var fieldDenylist = map[string]bool{
	"DB":      true,
	"Queries": true,
	"Tx":      true,
	"Ctx":     true,
	"Req":     true,
	"Resp":    true,
	"Server":  true,
}

// fieldIsDDLDrift tests whether the selector (`recv.Field` form) refers
// to a struct field the SSOT no longer provides. The receiver portion
// is discarded because the contract extractor records only exported
// selector names and the DDL Ground tracks column (field) identifiers
// alone. Returns true when the suffix after the final dot is not
// present in expected and is not a generator-internal accessor name.
//
// Selectors without a dot (shouldn't happen in practice — DDLFields
// entries are always "recv.Field") are treated as non-drift to avoid
// spurious diagnostics.
func fieldIsDDLDrift(selector string, expected map[string]bool) bool {
	idx := strings.LastIndex(selector, ".")
	if idx < 0 || idx == len(selector)-1 {
		return false
	}
	field := selector[idx+1:]
	if fieldDenylist[field] {
		return false
	}
	return !expected[canonicalFieldKey(field)]
}
