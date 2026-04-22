//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-14 — @state DiagramID 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s14StateDiagramID validates S-14.
func s14StateDiagramID(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "state" {
				continue
			}
			if seq.DiagramID == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-14] @state requires DiagramID",
					Advice:  "@state 시퀀스에 DiagramID 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
