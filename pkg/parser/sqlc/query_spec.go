//ff:type feature=orchestrator type=model
//ff:what sqlc query 한 건의 구조화된 스펙 (파일명→모델, name 주석→메서드)
package sqlc

// QuerySpec describes a single sqlc `-- name:` query entry.
//
// Name        : raw query name from the comment (e.g. "UserFindByEmail").
// Model       : PascalCase model derived from the filename (singularized).
// Method      : method name after stripping the model prefix. Falls back to
//               the full query name when the prefix does not match.
// Cardinality : ":one" / ":many" / ":exec" value without the leading colon.
// RowType     : sqlc-synthesized row struct name (e.g. "UserFindByEmailRow").
//               Populated for ":one" / ":many"; empty for ":exec" / ":execresult"
//               because those return no rows.
// File        : absolute path to the .sql file.
// Line        : 1-based line number of the `-- name:` comment.
type QuerySpec struct {
	Name        string
	Model       string
	Method      string
	Cardinality string
	RowType     string
	Params      []string // named params from SQL (@param_name → PascalCase "ParamName")
	File        string
	Line        int
}
