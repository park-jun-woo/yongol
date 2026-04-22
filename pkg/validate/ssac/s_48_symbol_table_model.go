//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-48 — the Model in @get/@post/@put/@delete must exist in the symbol table

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s48SymbolTableModel validates S-48: Model referenced by @get/@post/@put/@delete
// must appear in the loaded symbol table (Ground.Lookup["SymbolTable.model"]).
func s48SymbolTableModel(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	models, ok := g.Lookup["SymbolTable.model"]
	if !ok {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if !crudType(seq) {
				continue
			}
			model := extractModel(seq)
			if model == "" {
				continue
			}
			if models[model] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-48] model %q not found in symbol table", model),
				Advice:  fmt.Sprintf("Add or import the .ssac file that defines model %q", model),
			})
		}
	}
	return diags
}
