//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what jsxCompareOp — 가드 비교 연산자를 JS 연산자로 매핑 검증 (= → ===, != → !==, 나머지 유지)

package stml

import "testing"

func TestJsxCompareOp(t *testing.T) {
	tests := []struct {
		name string
		op   string
		want string
	}{
		{name: "equality", op: "=", want: "==="},
		{name: "inequality", op: "!=", want: "!=="},
		{name: "greater", op: ">", want: ">"},
		{name: "less", op: "<", want: "<"},
		{name: "greater equal", op: ">=", want: ">="},
		{name: "less equal", op: "<=", want: "<="},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jsxCompareOp(tt.op)
			if got != tt.want {
				t.Errorf("jsxCompareOp(%q) = %q, want %q", tt.op, got, tt.want)
			}
		})
	}
}
