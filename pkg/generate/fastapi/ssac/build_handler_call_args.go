//ff:func feature=gen-fastapi type=util control=sequence
//ff:what buildHandlerCallArgs — service 함수 호출 인자 목록(session/path/body/query/user/event_bus) 구성

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// buildHandlerCallArgs assembles the positional argument list passed from the
// route handler to the service function.
func buildHandlerCallArgs(plan *ir.ServicePlan, hasBody, isPreAuth, needsEventBus bool) []string {
	callArgs := []string{"session"}
	callArgs = append(callArgs, plan.PathParams...)
	if hasBody {
		callArgs = append(callArgs, "body")
	}
	callArgs = append(callArgs, queryParamNames(plan.QueryParams)...)
	if !isPreAuth {
		callArgs = append(callArgs, "current_user")
	}
	if needsEventBus {
		callArgs = append(callArgs, "event_bus")
	}
	return callArgs
}
