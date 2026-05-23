//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what collectPathVars — 빈 경로/단일/복수 변수 추출 검증

package openapi

import (
	"testing"
)

func TestCollectPathVars(t *testing.T) {
	tests := []struct {
		name string
		path string
		want map[string]bool
	}{
		{
			name: "no variables",
			path: "/users",
			want: map[string]bool{},
		},
		{
			name: "single variable",
			path: "/users/{id}",
			want: map[string]bool{"id": true},
		},
		{
			name: "multiple variables",
			path: "/orgs/{org_id}/users/{user_id}",
			want: map[string]bool{"org_id": true, "user_id": true},
		},
		{
			name: "empty string",
			path: "",
			want: map[string]bool{},
		},
		{
			name: "nested path with single var",
			path: "/api/v1/items/{item_id}/comments",
			want: map[string]bool{"item_id": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertPathVars(t, tt.path, tt.want)
		})
	}
}
