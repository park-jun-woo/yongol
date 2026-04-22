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
func (g *methodGen) buildFieldResponse(fields map[string]string) []string {
	var lines []string
	lines = append(lines, fmt.Sprintf("return api.%s200JSONResponse{", g.FuncName))

	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, jsonName := range keys {
		lines = append(lines, g.renderResponseField(jsonName, fields[jsonName]))
	}
	lines = append(lines, "}, nil")
	return lines
}
