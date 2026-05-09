//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.buildFieldResponse — 필드 맵을 typed response (db→api 변환 포함) 로 렌더링

package ssac

import (
	"fmt"
	"sort"
)

// buildFieldResponse generates typed response with db→api conversion.
// SSaC fields: { workflow: updated, count: cc.Count }
// OpenAPI RespFields tells us the api type for each field.
//
// convert<Model> now returns (*api.Model, error) so $ref-typed fields
// have their conversion hoisted to local variables before the struct
// literal and any error is propagated as nil,err from the handler
// (BUG-003 / BUG-005 response direction).
func (g *methodGen) buildFieldResponse(fields map[string]string) []string {
	var lines []string

	// 필드 값을 Go 표현식으로 변환 (request.id → request.Id 등)
	mapped := make(map[string]string, len(fields))
	for k, v := range fields {
		mapped[k] = g.mapValue(v)
	}

	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Pre-pass: hoist any $ref convert<Type>/convert<Type>List call into
	// a local variable so the per-call error is reachable. scalarLocal
	// keyed by jsonName → local var name; same for listLocal.
	scalarLocal := make(map[string]string)
	listLocal := make(map[string]string)
	for _, jsonName := range keys {
		rf, ok := g.RespFields[jsonName]
		if !ok || rf.RefType == "" {
			continue
		}
		varExpr := mapped[jsonName]
		local := lowerFirst(pascalCase(jsonName)) + "Converted"
		if rf.IsArray {
			listLocal[jsonName] = local
			lines = append(lines,
				fmt.Sprintf("%s, err := convert%sList(%s)", local, rf.RefType, varExpr),
				"if err != nil { return nil, err }",
			)
		} else {
			scalarLocal[jsonName] = local
			lines = append(lines,
				fmt.Sprintf("%s, err := convert%s(%s)", local, rf.RefType, varExpr),
				"if err != nil { return nil, err }",
			)
		}
	}

	// g.SuccessStatus is HTTP-method-derived at extract time (BUG-004) —
	// POST → 201, DELETE → 204, etc. instead of the prior hardcoded 200.
	lines = append(lines, fmt.Sprintf("return api.%s%dJSONResponse{",
		g.FuncName, g.SuccessStatus))

	for _, jsonName := range keys {
		lines = append(lines, g.renderResponseFieldHoisted(jsonName, mapped[jsonName], scalarLocal, listLocal))
	}
	lines = append(lines, "}, nil")
	return lines
}
