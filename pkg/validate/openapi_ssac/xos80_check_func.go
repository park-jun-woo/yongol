//ff:func feature=validate type=util control=sequence topic=ssac-openapi
//ff:what XOS-80 단일 ServiceFunc 검사 — Q4 회피용 추출 헬퍼

package openapi_ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	yopenapi "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xos80CheckFunc evaluates XOS-80 for a single SSaC service func and
// returns zero or one diagnostic. Extracted from xos80SuccessStatusMismatch
// to keep the range body under the Q4 pure-line budget.
func xos80CheckFunc(fn ssacparser.ServiceFunc, opMap map[string]OperationEntry) []diagnostic.Diagnostic {
	if !hasResponseSequence(fn) {
		return nil
	}
	entry, ok := opMap[fn.Name]
	if !ok {
		return nil
	}
	if yopenapi.DeriveSuccessStatus(entry.Op, entry.Method) != 0 {
		return nil
	}
	declared := sortedKeys(yopenapi.Declared2xx(entry.Op))
	return []diagnostic.Diagnostic{{
		File:  fn.FileName,
		Line:  fn.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XOS-80] operation %s (%s) cannot derive a success status — declared 2xx: %v",
			fn.Name, entry.Method, declared),
		Advice: "Declare a conventional 2xx response for this method (POST→201, PUT/PATCH→200, DELETE→204, GET→200)",
	}}
}
