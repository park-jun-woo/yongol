//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용

package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func bnFS(mc *pmanifest.ProjectConfig, funcs []ssac.ServiceFunc) *yongol.Fullstack {
	return &yongol.Fullstack{Manifest: mc, ServiceFuncs: funcs}
}

func bnCallFunc(prefix string) ssac.ServiceFunc {
	return ssac.ServiceFunc{
		Name:      "f",
		Sequences: []ssac.Sequence{{Type: "call", Model: prefix + "Put"}},
	}
}

func TestManifestDeclaresCache_ZeroCov(t *testing.T) {
	if manifestDeclaresCache(nil) {
		t.Error("nil should be false")
	}
	mc := &pmanifest.ProjectConfig{Cache: &pmanifest.BuiltinBackend{Backend: "redis"}}
	if !manifestDeclaresCache(bnFS(mc, nil)) {
		t.Error("declared cache should be true")
	}
}

func TestManifestDeclaresFile_ZeroCov(t *testing.T) {
	if manifestDeclaresFile(nil) {
		t.Error("nil should be false")
	}
	mc := &pmanifest.ProjectConfig{File: &pmanifest.FileBackend{Backend: "s3"}}
	if !manifestDeclaresFile(bnFS(mc, nil)) {
		t.Error("declared file should be true")
	}
}

func TestManifestDeclaresQueue_ZeroCov(t *testing.T) {
	if manifestDeclaresQueue(nil) {
		t.Error("nil should be false")
	}
	mc := &pmanifest.ProjectConfig{Queue: &pmanifest.QueueBackend{Backend: "postgres"}}
	if !manifestDeclaresQueue(bnFS(mc, nil)) {
		t.Error("declared queue should be true")
	}
}

func TestManifestDeclaresSession_ZeroCov(t *testing.T) {
	if manifestDeclaresSession(nil) {
		t.Error("nil should be false")
	}
	mc := &pmanifest.ProjectConfig{Session: &pmanifest.BuiltinBackend{Backend: "memory"}}
	if !manifestDeclaresSession(bnFS(mc, nil)) {
		t.Error("declared session should be true")
	}
}

func TestActiveBackendsFor_ZeroCov(t *testing.T) {
	mc := &pmanifest.ProjectConfig{
		Session: &pmanifest.BuiltinBackend{Backend: "memory"},
		Cache:   &pmanifest.BuiltinBackend{Backend: "redis"},
		File:    &pmanifest.FileBackend{Backend: "s3"},
		Queue:   &pmanifest.QueueBackend{Backend: "postgres"},
	}
	ab := activeBackendsFor(bnFS(mc, nil))
	if ab.Session == nil || ab.Cache == nil || ab.File == nil || ab.Queue == nil {
		t.Errorf("expected all backends resolved: %#v", ab)
	}
}

func TestCacheBackendFor_ZeroCov(t *testing.T) {
	mc := &pmanifest.ProjectConfig{Cache: &pmanifest.BuiltinBackend{Backend: "redis"}}
	if c := cacheBackendFor(bnFS(mc, nil)); c == nil || c.Backend != "redis" {
		t.Errorf("declared: %#v", c)
	}
	if c := cacheBackendFor(bnFS(nil, []ssac.ServiceFunc{bnCallFunc("cache.")})); c == nil || c.Backend != "memory" {
		t.Errorf("ssac default: %#v", c)
	}
	if c := cacheBackendFor(bnFS(nil, nil)); c != nil {
		t.Errorf("unused should be nil")
	}
}

func TestFileBackendFor_ZeroCov(t *testing.T) {
	mc := &pmanifest.ProjectConfig{File: &pmanifest.FileBackend{Backend: "s3"}}
	if f := fileBackendFor(bnFS(mc, nil)); f == nil || f.Backend != "s3" {
		t.Errorf("declared: %#v", f)
	}
	if f := fileBackendFor(bnFS(nil, []ssac.ServiceFunc{bnCallFunc("file.")})); f == nil || f.Backend != "local" {
		t.Errorf("ssac default: %#v", f)
	}
	if f := fileBackendFor(bnFS(nil, nil)); f != nil {
		t.Errorf("unused should be nil")
	}
}

func TestQueueBackendFor_ZeroCov(t *testing.T) {
	mc := &pmanifest.ProjectConfig{Queue: &pmanifest.QueueBackend{Backend: "postgres"}}
	if q := queueBackendFor(bnFS(mc, nil)); q == nil || q.Backend != "postgres" {
		t.Errorf("declared: %#v", q)
	}
	pub := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "publish"}}}
	if q := queueBackendFor(bnFS(nil, []ssac.ServiceFunc{pub})); q == nil || q.Backend != "postgres" {
		t.Errorf("ssac default: %#v", q)
	}
	if q := queueBackendFor(bnFS(nil, nil)); q != nil {
		t.Errorf("unused should be nil")
	}
}

func TestMiddlewaresFor_ZeroCov(t *testing.T) {
	if middlewaresFor(nil) != nil {
		t.Error("nil should be nil")
	}
	mc := &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Middleware: []string{"cors", "auth"}}}
	mws := middlewaresFor(bnFS(mc, nil))
	if len(mws) != 2 || mws[0].Name != "cors" {
		t.Errorf("middlewares wrong: %#v", mws)
	}
}

func TestRoutesFor_ZeroCov(t *testing.T) {
	if routesFor(bnFS(nil, nil)) != nil {
		t.Error("placeholder should be nil")
	}
}

func TestSequencesCallPrefix_ZeroCov(t *testing.T) {
	seqs := []ssac.Sequence{{Type: "call", Model: "session.Put"}, {Type: "get"}}
	if !sequencesCallPrefix(seqs, "session.") {
		t.Error("expected prefix match")
	}
	if sequencesCallPrefix(seqs, "cache.") {
		t.Error("unexpected match")
	}
}

func TestSsacUsesCacheCalls_ZeroCov(t *testing.T) {
	if ssacUsesCacheCalls(nil) {
		t.Error("nil false")
	}
	if !ssacUsesCacheCalls(bnFS(nil, []ssac.ServiceFunc{bnCallFunc("cache.")})) {
		t.Error("expected cache use")
	}
}

func TestSsacUsesFileCalls_ZeroCov(t *testing.T) {
	if ssacUsesFileCalls(nil) {
		t.Error("nil false")
	}
	if !ssacUsesFileCalls(bnFS(nil, []ssac.ServiceFunc{bnCallFunc("file.")})) {
		t.Error("expected file use")
	}
}

func TestSsacUsesQueue_ZeroCov(t *testing.T) {
	if ssacUsesQueue(nil) {
		t.Error("nil false")
	}
	pub := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "publish"}}}
	if !ssacUsesQueue(bnFS(nil, []ssac.ServiceFunc{pub})) {
		t.Error("expected queue use")
	}
}

func TestSsacUsesSessionCalls_ZeroCov(t *testing.T) {
	if ssacUsesSessionCalls(nil) {
		t.Error("nil false")
	}
	if !ssacUsesSessionCalls(bnFS(nil, []ssac.ServiceFunc{bnCallFunc("session.")})) {
		t.Error("expected session use")
	}
}

func TestAssertPrepareSessionBackend_ZeroCov(t *testing.T) {
	assertPrepareSessionBackend(t, prepareSessionBackendCase{wantNil: true}, nil)
	assertPrepareSessionBackend(t, prepareSessionBackendCase{wantBE: "memory"}, &Session{Backend: "memory"})
}
