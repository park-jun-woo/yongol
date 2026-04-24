//ff:func feature=validate type=rule control=sequence topic=states
//ff:what xsm27DiagForOperation — 단일 operation 에 대한 XSM-27 진단 빌드 (조건 미충족 시 (_, false))

package ssac_statemachine

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xsm27DiagForOperation evaluates a single (method, operation) pair and
// returns a XSM-27 diagnostic along with ok=true when all gate conditions
// hold. When any precondition misses (op missing, SSaC func absent,
// @state-neutral declared, @state already present, resource not read via
// FindByID) it returns the zero diagnostic with ok=false.
func xsm27DiagForOperation(
	_ string,
	op *openapi3.Operation,
	target *statefulTarget,
	funcByName map[string]ssac.ServiceFunc,
) (diagnostic.Diagnostic, bool) {
	if op == nil || op.OperationID == "" {
		return diagnostic.Diagnostic{}, false
	}
	fn, ok := funcByName[op.OperationID]
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	if fn.StateNeutral {
		return diagnostic.Diagnostic{}, false
	}
	if hasStateSequence(fn.Sequences) {
		return diagnostic.Diagnostic{}, false
	}
	resultVar, ok := findByIDResultVar(fn.Sequences, target.Model)
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	return buildXsm27Diag(fn, target, resultVar), true
}
