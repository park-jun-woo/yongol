//ff:func feature=validate type=rule control=iteration dimension=2 topic=sqlc
//ff:what XQS-14 — SSaC input key case ↔ sqlc param

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs14InputKeyCase validates XQS-14: SSaC input key case ↔ sqlc param.
// Input keys (from @get/@post/@put/@delete Args or @state/@auth/@publish
// Inputs) must match sqlc parameter names (DDL column names) case-sensitively.
// A case-insensitive-only match yields a WARNING.
func xqs14InputKeyCase(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil || g.Lookup == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			diags = append(diags, checkSeqInputKeyCase(fn, seq, g)...)
		}
	}
	return diags
}

