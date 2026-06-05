//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestRenameSourceVar -- dot-notation source 변수부 rename 치환(있/없음, dot 유/무) 검증

package ir

import "testing"

func TestRenameSourceVar(t *testing.T) {
	renames := map[string]string{"user": "user_result"}

	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"no dot, renamed", "user", "user_result"},
		{"no dot, not renamed", "params", "params"},
		{"dot, var renamed, suffix preserved", "user.email", "user_result.email"},
		{"dot, var not renamed", "body.title", "body.title"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renameSourceVar(tt.source, renames); got != tt.want {
				t.Errorf("renameSourceVar(%q) = %q, want %q", tt.source, got, tt.want)
			}
		})
	}
}
