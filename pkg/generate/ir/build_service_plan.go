//ff:func feature=gen-ir type=util control=iteration dimension=2
//ff:what BuildServicePlan -- SSaC ServiceFunc → ServicePlan IR 변환 (프레임워크 비의존, OpenAPI/DDL/Rego/StateDiagram 해석 정보 이식)

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

	// Extract OpenAPI metadata (path/query/body classification, success status).
	pathParams, queryParams := extractOpenAPIParams(fs, sf.Name, plan)

	// Build ops from sequences.
	for _, seq := range sf.Sequences {
		op, err := convertSequence(seq, fs)
		if err != nil {
			return nil, fmt.Errorf("BuildServicePlan(%s): %w", sf.Name, err)
		}
		plan.Ops = append(plan.Ops, op)
	}

	// Enrich FieldArg.Location for all ops using OpenAPI param classification.
	enrichFieldArgLocations(plan.Ops, pathParams, queryParams)

	// Enrich FieldArg.ColumnName and IsPK from DDL metadata.
	enrichFieldArgDDL(plan.Ops, fs)

	// Resolve variable shadowing: detect duplicate VarName declarations and
	// rename collisions with _result suffix. Reserve common method parameter
	// names so Op results do not shadow them (e.g. VerifyPassword ResultVar
	// "user" must not collide with the "user" method parameter).
	resolveVariableShadowing(plan.Ops, "params", "body", "query", "user", "payload")

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
