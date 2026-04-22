//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what diagram 이벤트 중 @state 가드가 없는 함수 수집

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// collectMissingGuards returns warnings for events that have a matching SSaC
// function but no @state sequence declared. funcByName maps function name to
// the ServiceFunc so the diagnostic can be anchored to the SSaC source file.
func collectMissingGuards(diagramID string, events []string, funcByName map[string]ssac.ServiceFunc, guardFuncs map[string]bool) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, event := range events {
		fn, ok := funcByName[event]
		if !ok || guardFuncs[event] {
			continue
		}
		file := fn.FileName
		if file == "" {
			file = "ssac/" + event + ".ssac"
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Line:    fn.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[XSM-26] function \"" + event + "\" participates in diagram \"" + diagramID + "\" but has no @state sequence",
			Advice:  "state 변경 전 @empty 가드를 추가하세요",
		})
	}
	return diags
}
