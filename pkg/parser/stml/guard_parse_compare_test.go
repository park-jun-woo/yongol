//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what parseGuardCompare — ref op value 비교식 파싱 (ident/string 값, 값 누락 에러) 검증

package stml

import "testing"

func TestParseGuardCompare(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantErr   bool
		wantOp    string
		wantValue string
	}{
		{name: "ident value", input: "a.b=active", wantOp: "=", wantValue: "active"},
		{name: "string value", input: "a.b='hi there'", wantOp: "=", wantValue: "hi there"},
		{name: "relational op", input: "a.b>=3", wantOp: ">=", wantValue: "3"},
		{name: "missing value", input: "a.b=&&", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertParseGuardCompare(t, tt.input, tt.wantErr, tt.wantOp, tt.wantValue)
		})
	}
}
