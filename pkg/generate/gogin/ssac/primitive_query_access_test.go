//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what primitiveQueryAccess 단위 테스트 (required면 accessor 그대로, 아니면 deref 래핑)

package ssac

import "testing"

func TestPrimitiveQueryAccess(t *testing.T) {
	cases := []struct {
		name     string
		required bool
		accessor string
		derefFn  string
		want     string
	}{
		{"required is passthrough", true, "request.Params.Limit", "derefInt", "request.Params.Limit"},
		{"optional int wrapped", false, "request.Params.Limit", "derefInt", "derefInt(request.Params.Limit)"},
		{"optional str wrapped", false, "request.Params.Q", "derefStr", "derefStr(request.Params.Q)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := primitiveQueryAccess(tc.required, tc.accessor, tc.derefFn)
			if got != tc.want {
				t.Errorf("primitiveQueryAccess = %q, want %q", got, tc.want)
			}
		})
	}
}
