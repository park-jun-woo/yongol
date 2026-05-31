//ff:func feature=manifest type=test control=sequence
//ff:what TestManifestHelpers — unit tests for the pure manifest parser helper functions
package manifest

import (
	"testing"
)

func TestParseRawClaims(t *testing.T) {
	raw := map[string]string{
		"OrgID":  "org_id:int64",
		"Email":  "email",
		"Region": "region:",
	}
	lines := map[string]int{"OrgID": 7}
	claims := parseRawClaims(raw, lines)

	org := claims["OrgID"]
	if org.Key != "org_id" || org.GoType != "int64" || !org.Typed || org.SourceLine != 7 {
		t.Errorf("OrgID = %+v", org)
	}
	email := claims["Email"]
	if email.Key != "email" || email.GoType != "string" || email.Typed {
		t.Errorf("Email = %+v", email)
	}
	// "region:" → empty type after colon defaults to string, not typed.
	region := claims["Region"]
	if region.Key != "region" || region.GoType != "string" || region.Typed {
		t.Errorf("Region = %+v", region)
	}

	// nil claimLines → SourceLine 0, no panic.
	c2 := parseRawClaims(map[string]string{"X": "x"}, nil)
	if c2["X"].SourceLine != 0 {
		t.Errorf("nil lines SourceLine = %d, want 0", c2["X"].SourceLine)
	}
}
