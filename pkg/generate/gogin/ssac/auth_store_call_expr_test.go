//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what authStoreCallExpr 단위 테스트 (RefreshRotate 4-arg vs 기타 3-arg 호출식)
package ssac

import (
	"testing"
)

func TestAuthStoreCallExpr(t *testing.T) {
	cases := []struct {
		name     string
		callFunc string
		want     string
	}{
		{"RefreshRotate is 4-arg", "RefreshRotate", "auth.RefreshRotate(ctx, nil, tok, false)"},
		{"Logout is 3-arg", "Logout", "auth.Logout(ctx, nil, tok)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := authStoreCallExpr("auth", tc.callFunc, "ctx", "tok")
			if got != tc.want {
				t.Errorf("authStoreCallExpr = %q, want %q", got, tc.want)
			}
		})
	}
}
