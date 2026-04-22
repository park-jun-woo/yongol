//ff:func feature=validate type=util control=iteration dimension=2 topic=sqlc
//ff:what checkSeqInputKeyCase — compare a single sequence's input keys against the sqlc param set for casing

package ssac_sqlc

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// checkSeqInputKeyCase compares a single sequence's input keys against the
// sqlc parameter set registered under "SQLc.param.<Model>".
func checkSeqInputKeyCase(fn ssac.ServiceFunc, seq ssac.Sequence, g *rule.Ground) []diagnostic.Diagnostic {
	if seq.Type == "call" || seq.Package != "" || seq.Model == "" {
		return nil
	}
	parts := strings.SplitN(seq.Model, ".", 2)
	if len(parts) < 2 {
		return nil
	}
	modelName := parts[0]
	params, ok := g.Lookup["SQLc.param."+modelName]
	if !ok || len(params) == 0 {
		return nil
	}
	keys := collectInputKeys(seq)
	if len(keys) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, key := range keys {
		if d, ok := checkSingleInputKeyCase(fn, seq, key, params); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
