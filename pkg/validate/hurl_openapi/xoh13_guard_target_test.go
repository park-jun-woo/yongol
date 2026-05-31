//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what TestHurlOpenAPIHelpers — unit tests for the pure hurl_openapi helper functions
package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXoh13GuardTarget(t *testing.T) {
	tests := []struct {
		name string
		seq  ssac.Sequence
		want string
	}{
		{"empty", ssac.Sequence{Type: "empty", Target: "course"}, "course"},
		{"exists", ssac.Sequence{Type: "exists", Target: "user.Email"}, "user.Email"},
		{"auth", ssac.Sequence{Type: "auth", Action: "delete", Resource: "project"}, "delete project"},
		{"state", ssac.Sequence{Type: "state", DiagramID: "reservation", Transition: "cancel"}, "reservation.cancel"},
		{"eval pkg", ssac.Sequence{Type: "eval", Package: "billing", Model: "Check"}, "billing.Check"},
		{"eval no pkg", ssac.Sequence{Type: "eval", Model: "Check"}, "Check"},
		{"other", ssac.Sequence{Type: "get"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xoh13GuardTarget(tt.seq); got != tt.want {
				t.Errorf("xoh13GuardTarget = %q, want %q", got, tt.want)
			}
		})
	}
}
