//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what XSS-59 — verifies that @subscribe message fields match the @publish payload

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xss59SubscribeFields validates XSS-59: every field declared on a @subscribe
// message struct must be supplied by some @publish payload on the same topic.
func xss59SubscribeFields(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	// Build topic -> set of payload field keys from every @publish.
	publishKeys := map[string]map[string]bool{}
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "publish" || seq.Topic == "" {
				continue
			}
			set := publishKeys[seq.Topic]
			if set == nil {
				set = map[string]bool{}
				publishKeys[seq.Topic] = set
			}
			for k := range seq.Inputs {
				set[k] = true
			}
			for k := range seq.Fields {
				set[k] = true
			}
		}
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe == nil || fn.Subscribe.Topic == "" || fn.Param == nil {
			continue
		}
		var msg *parsessac.StructInfo
		for i := range fn.Structs {
			if fn.Structs[i].Name == fn.Param.TypeName {
				msg = &fn.Structs[i]
				break
			}
		}
		if msg == nil {
			continue // S-39 reports missing struct
		}
		keys := publishKeys[fn.Subscribe.Topic]
		if keys == nil {
			continue // XSS-58 reports missing @publish
		}
		for _, field := range msg.Fields {
			if keys[field.Name] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    fn.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[XSS-59] @subscribe message field %q has no matching @publish payload key for topic %q", field.Name, fn.Subscribe.Topic),
				Advice:  fmt.Sprintf("Add the missing field %q to the @publish payload", field.Name),
			})
		}
	}
	return diags
}
