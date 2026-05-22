//ff:func feature=validate type=util control=sequence topic=ssac-openapi
//ff:what XOS-82 단일 ServiceFunc 검사 — Q1/Q4 회피용 추출 헬퍼

package openapi_ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	yopenapi "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xos82CheckFunc evaluates XOS-82 for a single SSaC service func and
// returns zero or one diagnostic. Extracted from
// xos82UnreachableSuccessStatus to keep nesting depth and Q4 pure-line
// count within limits.
func xos82CheckFunc(fn ssacparser.ServiceFunc, opMap map[string]OperationEntry) []diagnostic.Diagnostic {
	if !hasResponseSequence(fn) {
		return nil
	}
	entry, ok := opMap[fn.Name]
	if !ok {
		return nil
	}
	selected := yopenapi.DeriveSuccessStatus(entry.Op, entry.Method)
	if selected == 0 {
		return nil
	}
	declared := yopenapi.Declared2xx(entry.Op)
	if len(declared) <= 1 {
		return nil
	}
	unreachable := unreachable2xx(declared, selected)
	return []diagnostic.Diagnostic{{
		File:  fn.FileName,
		Line:  fn.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelWarning,
		Message: fmt.Sprintf(
			"[XOS-82] operation %s declares 2xx %v but only %d is reachable from SSaC",
			fn.Name, sortedKeys(declared), selected),
		Advice: fmt.Sprintf("Either remove the unused 2xx declarations %v or extend SSaC to emit them",
			sortedKeys(unreachable)),
		OperationID: fn.Name,
	}}
}
