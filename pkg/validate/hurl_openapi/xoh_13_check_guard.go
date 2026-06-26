//ff:func feature=validate type=rule control=sequence topic=hurl-openapi
//ff:what xoh13CheckGuard — guard 시퀀스 ErrStatus 의 hurl 커버리지 진단

package hurl_openapi

import (
	"fmt"
	"strconv"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func xoh13CheckGuard(fn ssac.ServiceFunc, seq ssac.Sequence, coveredSet map[string]bool, policies []rego.Policy) []diagnostic.Diagnostic {
	if !xoh13IsGuardType(seq.Type) {
		return nil
	}
	status := seq.ErrStatus
	if status == 0 {
		status = xoh13GuardDefaultStatus(seq.Type)
	}
	if status == 0 {
		return nil
	}
	if seq.Type == "auth" && isAllowAll(seq.Action, seq.Resource, policies) {
		return nil
	}
	statusStr := strconv.Itoa(status)
	if coveredSet != nil && coveredSet[statusStr] {
		return nil
	}

	target := xoh13GuardTarget(seq)
	return []diagnostic.Diagnostic{{
		File:    fn.FileName,
		Line:    seq.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: fmt.Sprintf("[XOH-13] %s — @%s %s \"%s\" %d has no hurl test", fn.Name, seq.Type, target, seq.Message, status),
		Advice:  fmt.Sprintf("Add a hurl scenario that triggers %s for this endpoint", statusStr),
	}}
}
