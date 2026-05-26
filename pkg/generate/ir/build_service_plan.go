//ff:func feature=gen-ir type=util control=iteration dimension=2
//ff:what BuildServicePlan -- SSaC ServiceFunc → ServicePlan IR 변환 (프레임워크 비의존)

package ir

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// BuildServicePlan converts an SSaC ServiceFunc into a framework-agnostic
// ServicePlan IR. The Fullstack context is used for cross-referencing OpenAPI
// and DDL metadata when needed. This function does NOT generate any backend
// code — it only produces the abstract execution plan.
func BuildServicePlan(sf *ssac.ServiceFunc, fs *yongol.Fullstack) (*ServicePlan, error) {
	plan := &ServicePlan{
		OperationID: sf.Name,
		FileName:    sf.FileName,
		Feature:     sf.Feature,
		Imports:     sf.Imports,
	}

	// Determine trigger kind.
	if sf.Subscribe != nil {
		plan.TriggerKind = TriggerSubscribe
		plan.Topic = sf.Subscribe.Topic
	} else {
		plan.TriggerKind = TriggerHTTP
	}

	// Build ops from sequences.
	for _, seq := range sf.Sequences {
		op, err := convertSequence(seq)
		if err != nil {
			return nil, fmt.Errorf("BuildServicePlan(%s): %w", sf.Name, err)
		}
		plan.Ops = append(plan.Ops, op)
	}

	// Annotate @get ops with lookahead guard info. When a @get is followed
	// by @empty or @exists targeting the same result variable, the renderer
	// must emit ErrNoRows tolerance instead of plain error propagation.
	annotateGetGuards(plan.Ops)

	// Determine transaction requirement.
	plan.UsesTransaction = planNeedsTransaction(plan.Ops)

	// Collect query methods.
	plan.QueryMethods = collectQueryMethods(sf.Sequences)

	return plan, nil
}
