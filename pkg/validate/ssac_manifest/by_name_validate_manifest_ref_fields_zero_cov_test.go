//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — manifest.* 참조 검증 + XNS-57 memory tx publish 직접 호출
package ssac_manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameValidateManifestRefFields_ZeroCov(t *testing.T) {
	mf := &pmanifest.ProjectConfig{}
	mf.Backend.Auth = &pmanifest.Auth{AccessTokenTTL: "15m"}
	fs := &yongol.Fullstack{Manifest: mf}

	// non-response sequence → nil.
	if d := validateManifestRefFields("Fn", "f.ssac", ssacparser.Sequence{Type: "post"}, fs); d != nil {
		t.Errorf("non-response should yield nil")
	}
	// response with a manifest.* field (unknown) → diag; non-manifest field skipped.
	seq := ssacparser.Sequence{Type: "response", Fields: map[string]string{
		"ttl":   "manifest.bogus",
		"other": "literal",
	}}
	if d := validateManifestRefFields("Fn", "f.ssac", seq, fs); len(d) != 1 {
		t.Errorf("expected 1 diag from unknown manifest ref, got %d", len(d))
	}
}
