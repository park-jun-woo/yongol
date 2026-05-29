//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what queryAccessExpr 단위 테스트 (enum/format/required 분기별 Go 접근식)

package ssac

import "testing"

func TestQueryAccessExpr(t *testing.T) {
	const acc = "request.Params.Sort"
	cases := []struct {
		name string
		qp   queryParam
		want string
	}{
		{"required enum cast to string", queryParam{IsEnum: true, IsRequired: true}, "string(request.Params.Sort)"},
		{"optional enum derefEnum", queryParam{IsEnum: true}, "derefEnum(request.Params.Sort)"},
		{"required integer passthrough", queryParam{GoType: "integer", IsRequired: true}, "request.Params.Sort"},
		{"optional integer derefInt", queryParam{GoType: "integer"}, "derefInt(request.Params.Sort)"},
		{"optional integer64 derefInt64", queryParam{GoType: "integer64"}, "derefInt64(request.Params.Sort)"},
		{"optional integer32 derefInt32", queryParam{GoType: "integer32"}, "derefInt32(request.Params.Sort)"},
		{"optional string derefStr", queryParam{GoType: "string"}, "derefStr(request.Params.Sort)"},
		{"optional boolean derefBool", queryParam{GoType: "boolean"}, "derefBool(request.Params.Sort)"},
		{"unknown defaults to derefStr", queryParam{GoType: "weird"}, "derefStr(request.Params.Sort)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := queryAccessExpr(tc.qp, acc); got != tc.want {
				t.Errorf("queryAccessExpr = %q, want %q", got, tc.want)
			}
		})
	}
}
