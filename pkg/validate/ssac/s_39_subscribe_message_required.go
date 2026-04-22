//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-39 — a struct matching the @subscribe message type must be defined

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s39SubscribeMessageRequired validates S-39: @subscribe Param's TypeName must
// have a matching Go struct definition in the same .ssac file.
func s39SubscribeMessageRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe == nil || fn.Param == nil || fn.Param.TypeName == "" {
			continue
		}
		found := false
		for _, s := range fn.Structs {
			if s.Name == fn.Param.TypeName {
				found = true
				break
			}
		}
		if found {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fn.FileName,
			Line:    fn.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[S-39] @subscribe message type %s has no struct definition", fn.Param.TypeName),
			Advice:  fmt.Sprintf("Define a struct corresponding to the @subscribe message type %s", fn.Param.TypeName),
		})
	}
	return diags
}
