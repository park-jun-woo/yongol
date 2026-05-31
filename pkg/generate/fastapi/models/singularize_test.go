//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestSingularize — 영어 복수형 snake_case → 단수형 변환 검증
package models

import (
	"testing"
)

func TestSingularize(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"workflows", "workflow"},
		{"categories", "category"},
		{"boxes", "box"},
		{"users", "user"},
		{"process", "process"},
		{"addresses", "address"},
		{"templates", "template"},
		{"execution_logs", "execution_log"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := singularize(tt.input)
			if got != tt.want {
				t.Errorf("singularize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
