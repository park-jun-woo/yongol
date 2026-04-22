//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what XSS-57 — verifies that every @publish topic has a corresponding @subscribe function

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xss57PublishToSubscribe validates XSS-57: every @publish topic must have a
// matching @subscribe function declared somewhere in the SSaC corpus.
func xss57PublishToSubscribe(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	subscribes := map[string]bool{}
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil && fn.Subscribe.Topic != "" {
			subscribes[fn.Subscribe.Topic] = true
		}
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "publish" || seq.Topic == "" {
				continue
			}
			if subscribes[seq.Topic] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[XSS-57] @publish topic %q has no matching @subscribe", seq.Topic),
				Advice:  fmt.Sprintf("Add a @subscribe function for topic %q", seq.Topic),
			})
		}
	}
	return diags
}
