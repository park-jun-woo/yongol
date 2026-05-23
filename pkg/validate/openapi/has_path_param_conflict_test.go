//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what hasPathParamConflict — 중복 없음/중복 있음/비변수 세그먼트 검증

package openapi

import "testing"

func TestHasPathParamConflict(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "no variables",
			path: "/users",
			want: false,
		},
		{
			name: "single variable no conflict",
			path: "/users/{id}",
			want: false,
		},
		{
			name: "different variables no conflict",
			path: "/orgs/{org_id}/users/{user_id}",
			want: false,
		},
		{
			name: "duplicate variable has conflict",
			path: "/orgs/{id}/users/{id}",
			want: true,
		},
		{
			name: "empty path",
			path: "",
			want: false,
		},
		{
			name: "non-variable duplicate segments no conflict",
			path: "/api/api/users",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasPathParamConflict(tt.path)
			if got != tt.want {
				t.Errorf("hasPathParamConflict(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}
