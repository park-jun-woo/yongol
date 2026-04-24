//ff:func feature=validate type=rule control=sequence topic=states
//ff:what buildXsm27Advice — XSM-27 advice 본문 문자열 조립

package ssac_statemachine

import (
	"strings"
)

// buildXsm27Advice returns the multi-line advice body for a XSM-27 WARNING.
// Extracted from buildXsm27Diag so each func stays inside the Q4 PURE
// line budget and Option A / self-loop-hint / Option B concerns are
// localised.
func buildXsm27Advice(fnName string, target *statefulTarget, diagramID, initial, varName string) string {
	var b strings.Builder
	writeXsm27OptionA(&b, fnName, target, diagramID, varName)
	if diagramID != "" && initial != "" {
		writeXsm27SelfLoopHint(&b, fnName, diagramID, initial)
	}
	writeXsm27OptionB(&b)
	return b.String()
}
