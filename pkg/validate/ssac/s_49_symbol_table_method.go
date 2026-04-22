//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-49 — Model.Method 중 Method가 심볼 테이블에 존재

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s49SymbolTableMethod validates S-49: Method portion of Model.Method must
// appear under Ground.Lookup["SymbolTable.method.<Model>"].
func s49SymbolTableMethod(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if !crudType(seq) {
				continue
			}
			model := extractModel(seq)
			method := extractMethod(seq)
			if model == "" || method == "" {
				continue
			}
			methods, ok := g.Lookup["SymbolTable.method."+model]
			if !ok {
				continue
			}
			if methods[method] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-49] method %s not found on model %s", method, model),
				Advice:  fmt.Sprintf("model %s 에 method %s 를 정의하세요 (예: @method %s Model.%s(...))", model, method, method, method),
			})
		}
	}
	return diags
}
