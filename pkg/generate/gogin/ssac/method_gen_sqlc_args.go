//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.sqlcArgs — Inputs를 sqlc 호출 인자 문자열로 변환 (0/1/N 분기)

package ssac

import (
	"sort"
	"strings"
)

// sqlcArgs converts Inputs → sqlc call arguments string.
// Keys are converted to sqlc PascalCase (e.g. "URL" → "Url", "OrgID" → "OrgID").
func (g *methodGen) sqlcArgs(method string, inputs map[string]string) string {
	if len(inputs) == 0 {
		return "ctx"
	}
	if len(inputs) == 1 {
		for _, v := range inputs {
			return "ctx, " + g.mapValue(v)
		}
	}
	// Validate (XQS-14/16) guarantees key matches sqlc field name. Use as-is.
	var fields []string
	for k, v := range inputs {
		fields = append(fields, k+": "+g.mapValue(v))
	}
	sort.Strings(fields)
	return "ctx, db." + method + "Params{" + strings.Join(fields, ", ") + "}"
}
