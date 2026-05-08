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

	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Pre-pass: hoist any $ref convert<Type>/convert<Type>List call into
	// a local variable so the per-call error is reachable. scalarLocal
	// keyed by jsonName → local var name; same for listLocal.
	// Skip converter for @call result variables — Func Response structs
	// are user-authored OpenAPI-compatible types (BUG-050).
	scalarLocal := make(map[string]string)
	listLocal := make(map[string]string)
	directAssign := make(map[string]bool)
	for _, jsonName := range keys {
		rf, ok := g.RespFields[jsonName]
		if !ok || rf.RefType == "" {
			continue
		}
		varExpr := fields[jsonName]
		if g.CallResultVars[varExpr] {
			directAssign[jsonName] = true
			continue
		}
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
		if directAssign[jsonName] {
			lines = append(lines, g.renderDirectAssignField(jsonName, fields[jsonName]))
		} else {
			lines = append(lines, g.renderResponseFieldHoisted(jsonName, fields[jsonName], scalarLocal, listLocal))
		}
	}
	lines = append(lines, "}, nil")
	return lines
}
