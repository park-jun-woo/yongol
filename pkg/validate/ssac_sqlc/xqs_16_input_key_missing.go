//ff:func feature=validate type=rule control=iteration dimension=3 topic=sqlc
//ff:what XQS-16 — ERROR when a SSaC Input key is absent from sqlc Params

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs16InputKeyMissing validates XQS-16: every SSaC CRUD Input key must exist
// as a sqlc Params field. Catches typos and name mismatches.
// Skip: seq.Type == "call", seq.Package != "".
func xqs16InputKeyMissing(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	paramMap := buildQueryParamMap(fs)
	if len(paramMap) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type == "call" || seq.Package != "" {
				continue
			}
			if seq.Model == "" || len(seq.Inputs) == 0 {
				continue
			}
			queryName := resolveQueryName(seq)
			params, ok := paramMap[queryName]
			if !ok {
				continue // S-49 handles missing query
			}
			for key := range seq.Inputs {
				if params[key] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[XQS-16] SSaC Input key %q is not in Params of sqlc query %s", key, queryName),
					Advice:  fmt.Sprintf("Add a @%s named parameter to the SQL query, or remove it from the SSaC Input", toSnake(key)),
				})
			}
		}
	}
	return diags
}
