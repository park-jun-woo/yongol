//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestUsesCurrentUser(t *testing.T) {
	// auth sequence.
	if !usesCurrentUser([]ssac.ServiceFunc{{Sequences: []ssac.Sequence{{Type: "auth"}}}}) {
		t.Error("auth seq should count")
	}
	// currentUser. input.
	withInput := []ssac.ServiceFunc{{Sequences: []ssac.Sequence{
		{Type: "post", Inputs: map[string]string{"owner": "currentUser.ID"}},
	}}}
	if !usesCurrentUser(withInput) {
		t.Error("currentUser. input should count")
	}
	// no usage.
	if usesCurrentUser([]ssac.ServiceFunc{{Sequences: []ssac.Sequence{
		{Type: "post", Inputs: map[string]string{"x": "body.X"}},
	}}}) {
		t.Error("plain input should not count")
	}
}
