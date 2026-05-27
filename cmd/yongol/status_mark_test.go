//ff:func feature=cli type=test control=iteration dimension=1
//ff:what TestStatusMark — statusMark 심볼 매핑 검증

package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestStatusMark(t *testing.T) {
	tests := []struct {
		name string
		st   validate.Status
		want string
	}{
		{"Pass", validate.StatusPass, "✓"},
		{"Fail", validate.StatusFail, "✗"},
		{"Skip", validate.StatusSkip, "-"},
		{"Missing", validate.StatusMissing, "?"},
		{"Unknown", validate.Status(99), " "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := statusMark(tt.st)
			if got != tt.want {
				t.Errorf("statusMark(%v) = %q, want %q", tt.st, got, tt.want)
			}
		})
	}
}
