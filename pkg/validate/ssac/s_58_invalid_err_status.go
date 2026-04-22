//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-58 — HTTP status code is not registered with IANA

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s58InvalidErrStatus validates S-58: ErrStatus must be a valid HTTP status (100-599).
func s58InvalidErrStatus(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.ErrStatus == 0 {
				continue
			}
			if seq.ErrStatus >= 100 && seq.ErrStatus <= 599 {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-58] ErrStatus %d is not a valid HTTP status code (100-599)", seq.ErrStatus),
				Advice:  "Use an IANA-registered HTTP status code (4xx/5xx)",
			})
		}
	}
	return diags
}
