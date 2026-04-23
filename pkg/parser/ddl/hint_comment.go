//ff:type feature=manifest type=model
//ff:what HintComment — DDL 주석 내 `-- @<tag> key=val ...` 힌트 레코드

package ddl

// HintComment is a single `-- @<tag> key=val ...` comment extracted from
// a DDL file, together with enough context for yongol migration hints
// to associate it with the right table/column.
//
// The scanner is intentionally simple: it records every hint comment
// and its *immediate context* — the CREATE TABLE currently being
// scanned (if any) plus the previous non-blank token on the same line
// (the candidate column name for column-line hints).
type HintComment struct {
	File       string            // absolute SQL file path
	Line       int               // 1-based
	Tag        string            // "rename" / "cast" / "backfill" / "data_migration" / "allow_destructive"
	Args       map[string]string // key=val pairs after the tag
	TableCtx   string            // last CREATE TABLE entered (lowercase), "" outside any table
	ColumnCtx  string            // column name if the hint is on a column definition line, "" otherwise
	BlockAbove bool              // true when the hint is on its own line *above* a CREATE TABLE / column
}
