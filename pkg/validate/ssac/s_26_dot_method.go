//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-26 — Model.Method 형식 필수

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s26DotMethod validates S-26: @get/@post/@put/@delete Model must be Model.Method.
func s26DotMethod(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if !crudType(seq) {
				continue
			}
			if seq.Model == "" {
				continue // S-1/3/6/9 already report
			}
			idx := strings.IndexByte(seq.Model, '.')
			if idx <= 0 || idx == len(seq.Model)-1 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-26] Model must be Model.Method format (got %q)", seq.Model),
					Advice:  "Model.Method 형식으로 작성하세요 (예: User.FindByEmail)",
				})
			}
		}
	}
	return diags
}
