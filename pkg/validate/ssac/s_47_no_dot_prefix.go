//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-47 — @get/@post/@put/@delete Model에 package prefix 금지

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s47NoDotPrefix validates S-47: CRUD Model must not carry a "package." prefix
// — Model.Method format only (one dot).
func s47NoDotPrefix(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if !crudType(seq) || seq.Model == "" {
				continue
			}
			if strings.Count(seq.Model, ".") <= 1 && seq.Package == "" {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-47] package-prefix @model not allowed (got %q)", seq.Model),
				Advice:  "@model 값에서 패키지 prefix(pkg.X) 를 제거하세요",
			})
		}
	}
	return diags
}
