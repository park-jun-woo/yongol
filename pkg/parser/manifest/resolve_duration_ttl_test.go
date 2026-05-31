//ff:func feature=manifest type=test control=sequence
//ff:what TestManifestHelpers — unit tests for the pure manifest parser helper functions
package manifest

import (
	"testing"
)

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
