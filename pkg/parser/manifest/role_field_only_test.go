//ff:func feature=projectconfig type=test control=iteration dimension=1
//ff:what TestRoleFieldOnly — role_field 전용 블록 판정의 양성·음성(nil/혼재 키) 검증

package manifest

import "testing"

func TestRoleFieldOnly(t *testing.T) {
	cases := []struct {
		name string
		auth *FrontendAuth
		want bool
	}{
		{"nil block", nil, false},
		{"empty block", &FrontendAuth{}, false},
		{"role_field only", &FrontendAuth{RoleField: "role"}, true},
		{"role_field + token_field", &FrontendAuth{RoleField: "role", TokenField: "access_token"}, false},
		{"role_field + refresh_field", &FrontendAuth{RoleField: "role", RefreshField: "refresh_token"}, false},
		{"role_field + refresh_op", &FrontendAuth{RoleField: "role", RefreshOp: "Refresh"}, false},
		{"role_field + store", &FrontendAuth{RoleField: "role", Store: "memory"}, false},
		{"token_field only", &FrontendAuth{TokenField: "access_token"}, false},
	}
	for _, c := range cases {
		if got := c.auth.RoleFieldOnly(); got != c.want {
			t.Errorf("%s: RoleFieldOnly() = %v, want %v", c.name, got, c.want)
		}
	}
}
