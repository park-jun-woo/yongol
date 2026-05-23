//ff:func feature=validate type=test control=iteration dimension=1 topic=response-body-required
//ff:what opIDOrPath — operationId 우선 + 빈 문자열 fallback 검증

package openapi

import "testing"

func TestOpIDOrPath(t *testing.T) {
	tests := []struct {
		name string
		opID string
		path string
		want string
	}{
		{"operationId present", "getUser", "/users/{id}", "getUser"},
		{"operationId empty", "", "/users/{id}", "/users/{id}"},
		{"both empty", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := opIDOrPath(tt.opID, tt.path)
			if got != tt.want {
				t.Errorf("opIDOrPath(%q, %q) = %q, want %q", tt.opID, tt.path, got, tt.want)
			}
		})
	}
}
