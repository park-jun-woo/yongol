//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what pathSegments — 빈 경로/변수 제외/역순 반환 검증

package openapi_ddl

import (
	"reflect"
	"testing"
)

func TestPathSegments(t *testing.T) {
	tests := []struct {
		name string
		path string
		want []string
	}{
		{"empty path", "", []string{}},
		{"root only", "/", []string{}},
		{"single segment", "/users", []string{"users"}},
		{"param skipped", "/users/{id}", []string{"users"}},
		{"multiple segments reversed", "/api/v1/users", []string{"users", "v1", "api"}},
		{"mixed segments and params", "/api/users/{id}/posts", []string{"posts", "users", "api"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathSegments(tt.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSegments(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}
