//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — manifest.* 참조 검증 + XNS-57 memory tx publish 직접 호출

package ssac_manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
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

func TestByNameXns57MemoryTxPublish_ZeroCov(t *testing.T) {
	// nil fs → nil.
	if d := xns57MemoryTxPublish(nil); d != nil {
		t.Errorf("nil fs should yield nil")
	}
	// backend != memory → nil.
	fsPg := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Queue: &pmanifest.QueueBackend{Backend: "postgres"}}}
	if d := xns57MemoryTxPublish(fsPg); d != nil {
		t.Errorf("postgres backend should yield nil, got %v", d)
	}
	// memory backend + tx-bound publish func → warning.
	fsMem := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Queue: &pmanifest.QueueBackend{Backend: "memory"}},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name:     "CreateOrder",
			FileName: "create_order.ssac",
			Sequences: []ssacparser.Sequence{
				{Type: "post", Topic: ""},
				{Type: "publish", Topic: "order.created"},
			},
		}},
	}
	d := xns57MemoryTxPublish(fsMem)
	if len(d) != 1 {
		t.Errorf("memory + tx-bound publish should yield 1 warning, got %d", len(d))
	}
}
