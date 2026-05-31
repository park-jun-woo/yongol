//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — manifest.* 참조 검증 + XNS-57 memory tx publish 직접 호출
package ssac_manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestByNameValidateOneManifestRef_ZeroCov(t *testing.T) {
	mf := &pmanifest.ProjectConfig{}
	mf.Backend.Auth = &pmanifest.Auth{AccessTokenTTL: "15m"}

	// unknown manifest path → XNS-80 diag.
	if d := validateOneManifestRef("Fn", "f.ssac", 1, "ttl", "manifest.bogus.path", mf); len(d) != 1 {
		t.Errorf("unknown path should yield 1 diag, got %d", len(d))
	}
	// known + present → nil.
	if d := validateOneManifestRef("Fn", "f.ssac", 1, "ttl", "manifest.auth.accessTokenTTL", mf); d != nil {
		t.Errorf("known present ref should yield nil, got %v", d)
	}
	// known but missing value → diag.
	mfEmpty := &pmanifest.ProjectConfig{}
	if d := validateOneManifestRef("Fn", "f.ssac", 1, "ttl", "manifest.auth.accessTokenTTL", mfEmpty); len(d) != 1 {
		t.Errorf("known-but-missing should yield 1 diag, got %d", len(d))
	}
}
