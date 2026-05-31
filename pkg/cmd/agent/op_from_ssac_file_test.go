//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestOpFromSSaCFile — SSaC 파일명/경로에서 operationId(.ssac 제거) 추출 검증
package agent

import (
	"testing"
)

func TestOpFromSSaCFile(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"service/user/Login.ssac", "Login"},
		{"Login.ssac", "Login"},
		{"service/order/CreateOrder.ssac", "CreateOrder"},
		{"NoExt", "NoExt"},
	}
	for _, c := range cases {
		if got := opFromSSaCFile(c.in); got != c.want {
			t.Errorf("opFromSSaCFile(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
