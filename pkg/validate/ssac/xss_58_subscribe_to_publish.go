//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what XSS-58 — verifies that every @subscribe topic has a corresponding @publish sequence

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xss58SubscribeToPublish validates XSS-58: every @subscribe topic must have a
// matching @publish sequence somewhere in the SSaC corpus.
func xss58SubscribeToPublish(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	publishes := map[string]bool{}
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type == "publish" && seq.Topic != "" {
				publishes[seq.Topic] = true
			}
		}
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe == nil || fn.Subscribe.Topic == "" {
			continue
		}
		if publishes[fn.Subscribe.Topic] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fn.FileName,
			Line:    fn.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XSS-58] @subscribe topic %q has no matching @publish", fn.Subscribe.Topic),
			Advice:  fmt.Sprintf("Add a @publish sequence for topic %q", fn.Subscribe.Topic),
		})
	}
	return diags
}
