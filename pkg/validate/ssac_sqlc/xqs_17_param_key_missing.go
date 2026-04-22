//ff:func feature=validate type=rule control=iteration dimension=3 topic=sqlc
//ff:what XQS-17 — ERROR when a sqlc Params field is absent from the SSaC Input

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs17ParamKeyMissing validates XQS-17: every sqlc Params field must be
// referenced by at least one SSaC CRUD Input key. Catches forgotten params.
// Skip: seq.Type == "call", seq.Package != "".
func xqs17ParamKeyMissing(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
			if seq.Model == "" {
				continue
			}
			queryName := resolveQueryName(seq)
			params, ok := paramMap[queryName]
			if !ok {
				continue
			}
			inputKeys := make(map[string]bool)
			for k := range seq.Inputs {
				inputKeys[k] = true
			}
			for param := range params {
				if inputKeys[param] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[XQS-17] sqlc query %s Params field %q is missing from the SSaC Input", queryName, param),
					Advice:  fmt.Sprintf("Add {%s: <value>} to the SSaC @%s Inputs", param, seq.Type),
				})
			}
		}
	}
	return diags
}
