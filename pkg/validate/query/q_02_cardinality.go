//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-02 — cardinality (:one / :many / :exec / :execrows) is required

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

var validCardinalities = map[string]bool{
	"one":      true,
	"many":     true,
	"exec":     true,
	"execrows": true,
	"execlastid": true,
}

// q02Cardinality validates Q-02: every `-- name:` must declare a valid
// cardinality. sqlc itself would reject missing cardinality at generate
// time, but catching it here gives immediate feedback.
func q02Cardinality(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, q := range fs.SQLcQueries {
		if validCardinalities[q.Cardinality] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    q.File,
			Line:    q.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[Q-02] query " + q.Name + " has missing or invalid cardinality (" + q.Cardinality + ")",
			Advice:  "Specify one of :one / :many / :exec / :execrows, e.g. `-- name: " + q.Name + " :one`",
		})
	}
	return diags
}
