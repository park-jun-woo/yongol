//ff:func feature=manifest type=test control=sequence
//ff:what TestManifestHelpers — unit tests for the pure manifest parser helper functions
package manifest

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestIntToStr(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{0, "0"},
		{42, "42"},
		{-7, "-7"},
		{900, "900"},
		{-1000000, "-1000000"},
	}
	for _, tt := range tests {
		if got := intToStr(tt.in); got != tt.want {
			t.Errorf("intToStr(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestResolveDurationTTL(t *testing.T) {
	access := func(a *Auth) string { return a.AccessTokenTTL }

	auth := &Auth{AccessTokenTTL: "15m"}
	rv, ok := resolveDurationTTL(auth, access)
	if !ok {
		t.Fatal("expected ok for 15m")
	}
	if rv.GoLit != "900" || rv.GoType != "int64" || rv.Raw != "15m" {
		t.Errorf("rv = %+v, want {Raw:15m GoLit:900 GoType:int64}", rv)
	}

	// nil auth → not ok.
	if _, ok := resolveDurationTTL(nil, access); ok {
		t.Error("nil auth should not resolve")
	}
	// empty value → not ok.
	if _, ok := resolveDurationTTL(&Auth{}, access); ok {
		t.Error("empty TTL should not resolve")
	}
	// unparseable → not ok.
	if _, ok := resolveDurationTTL(&Auth{AccessTokenTTL: "notaduration"}, access); ok {
		t.Error("invalid duration should not resolve")
	}
}

func TestResolvedMode(t *testing.T) {
	if got := (&Auth{}).ResolvedMode(); got != "cookie" {
		t.Errorf("empty mode → %q, want cookie", got)
	}
	if got := (&Auth{Mode: "bearer"}).ResolvedMode(); got != "bearer" {
		t.Errorf("explicit bearer → %q", got)
	}
	var nilAuth *Auth
	if got := nilAuth.ResolvedMode(); got != "cookie" {
		t.Errorf("nil auth → %q, want cookie", got)
	}
}

func TestMappingValue(t *testing.T) {
	var root yaml.Node
	if err := yaml.Unmarshal([]byte("a: 1\nb: hello\n"), &root); err != nil {
		t.Fatal(err)
	}
	doc := root.Content[0] // document → mapping node
	if v := mappingValue(doc, "b"); v == nil || v.Value != "hello" {
		t.Errorf("mappingValue(b) = %v", v)
	}
	if v := mappingValue(doc, "missing"); v != nil {
		t.Errorf("missing key should return nil, got %v", v)
	}
	// nil node → nil.
	if v := mappingValue(nil, "a"); v != nil {
		t.Error("nil node → nil")
	}
	// non-mapping node → nil.
	scalar := &yaml.Node{Kind: yaml.ScalarNode, Value: "x"}
	if v := mappingValue(scalar, "a"); v != nil {
		t.Error("scalar node → nil")
	}
}

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
