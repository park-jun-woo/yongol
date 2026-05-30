//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestHurlOpenAPIHelpers — unit tests for the pure hurl_openapi helper functions
package hurl_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectNon5xxCodes(t *testing.T) {
	got := collectNon5xxCodes(map[string]bool{"200": true, "404": true, "500": true, "503": true})
	want := []string{"200", "404"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("collectNon5xxCodes = %v, want %v", got, want)
	}
	if got := collectNon5xxCodes(map[string]bool{"500": true}); got != nil {
		t.Errorf("only-5xx → %v, want nil", got)
	}
}

func TestFilterUncovered(t *testing.T) {
	declared := []string{"200", "404", "409"}
	covered := map[string]bool{"200": true}
	got := filterUncovered(declared, covered)
	want := []string{"404", "409"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filterUncovered = %v, want %v", got, want)
	}
	// nil coveredSet → all declared are uncovered.
	if got := filterUncovered(declared, nil); !reflect.DeepEqual(got, declared) {
		t.Errorf("nil covered → %v, want %v", got, declared)
	}
}

func TestFormatCoveredCodes(t *testing.T) {
	if got := formatCoveredCodes(map[string]bool{"201": true, "200": true}); got != "200, 201" {
		t.Errorf("got %q, want '200, 201'", got)
	}
	if got := formatCoveredCodes(map[string]bool{}); got != "none" {
		t.Errorf("empty → %q, want none", got)
	}
}

func TestXoh13IsGuardType(t *testing.T) {
	for _, ty := range []string{"empty", "exists", "auth", "state", "eval"} {
		if !xoh13IsGuardType(ty) {
			t.Errorf("%q should be a guard type", ty)
		}
	}
	for _, ty := range []string{"get", "post", "publish", ""} {
		if xoh13IsGuardType(ty) {
			t.Errorf("%q should not be a guard type", ty)
		}
	}
}

func TestXoh13GuardDefaultStatus(t *testing.T) {
	tests := []struct {
		ty   string
		want int
	}{
		{"empty", 404},
		{"exists", 409},
		{"auth", 403},
		{"state", 409},
		{"eval", 0},
		{"get", 0},
	}
	for _, tt := range tests {
		if got := xoh13GuardDefaultStatus(tt.ty); got != tt.want {
			t.Errorf("xoh13GuardDefaultStatus(%q) = %d, want %d", tt.ty, got, tt.want)
		}
	}
}

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
