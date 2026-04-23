//ff:type feature=report type=model topic=json
//ff:what Summary — 전체 실행 카운트 집계 (errors / warnings / checks)
package json

// Summary aggregates counts across the whole run.
//
// Checks is the number of validation rules executed in this run (i.e. the
// embedded catalog size, not the rules that fired). Consumers can use it as
// a sanity check — "0 checks" means the catalog never loaded.
type Summary struct {
	Errors   int `json:"errors"`
	Warnings int `json:"warnings"`
	Checks   int `json:"checks"`
}
