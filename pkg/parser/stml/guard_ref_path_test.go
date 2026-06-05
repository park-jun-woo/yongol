//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what GuardRef.Path — model.field dotted 문자열 반환 검증

package stml

import "testing"

func TestGuardRefPath(t *testing.T) {
	tests := []struct {
		name string
		ref  GuardRef
		want string
	}{
		{name: "typical ref", ref: GuardRef{Model: "workflow", Field: "status"}, want: "workflow.status"},
		{name: "empty parts", ref: GuardRef{}, want: "."},
		{name: "model only", ref: GuardRef{Model: "user"}, want: "user."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ref.Path()
			if got != tt.want {
				t.Errorf("Path() = %q, want %q", got, tt.want)
			}
		})
	}
}
