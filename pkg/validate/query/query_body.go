//ff:type feature=validate type=model topic=query-structural
//ff:what queryBody — 단일 sqlc 쿼리 본문(헤더/라인/텍스트/escape flags)

package query

// queryBody is the raw text of a single sqlc query, starting at the
// `-- name:` comment and ending at the next `-- name:` (or EOF).
type queryBody struct {
	Header  string   // the `-- name:` line
	Lines   []string // body lines after the header, until next header / EOF
	Text    string   // full body joined
	HasStop bool     // has `-- @no-pagination`/`-- @allow-truncate`/`-- @allow-sensitive` etc.
	Escapes map[string]bool
}
