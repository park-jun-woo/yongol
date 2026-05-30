//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XSA-71/72/74-77 + Run test — backend required(ERROR)/unused(WARNING) 규칙 검증

package ssac_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func publishFunc() ssac.ServiceFunc {
	return ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "publish"}}}
}

func TestXsa71CacheBackendRequired(t *testing.T) {
	if d := xsa71CacheBackendRequired(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	// Not using cache -> nil.
	if d := xsa71CacheBackendRequired(fsWithFuncs()); d != nil {
		t.Errorf("no cache use -> nil, got %v", d)
	}
	usesCacheFn := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	// Backend declared -> nil.
	usesCacheFn.Manifest = &manifest.ProjectConfig{Cache: &manifest.BuiltinBackend{Backend: "redis"}}
	if d := xsa71CacheBackendRequired(usesCacheFn); d != nil {
		t.Errorf("declared backend -> nil, got %v", d)
	}
	// Backend missing -> ERROR.
	missing := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	d := xsa71CacheBackendRequired(missing)
	if len(d) != 1 || d[0].Level != diagnostic.LevelError || !strings.Contains(d[0].Message, "[XSA-71]") {
		t.Fatalf("expected XSA-71 error, got %v", d)
	}
}

func TestXsa72FileBackendRequired(t *testing.T) {
	if d := xsa72FileBackendRequired(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	usesFileFn := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("file.Save")}})
	usesFileFn.Manifest = &manifest.ProjectConfig{File: &manifest.FileBackend{Backend: "local"}}
	if d := xsa72FileBackendRequired(usesFileFn); d != nil {
		t.Errorf("declared backend -> nil, got %v", d)
	}
	missing := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("file.Save")}})
	d := xsa72FileBackendRequired(missing)
	if len(d) != 1 || d[0].Level != diagnostic.LevelError || !strings.Contains(d[0].Message, "[XSA-72]") {
		t.Fatalf("expected XSA-72 error, got %v", d)
	}
}

func TestXsa74SessionBackendUnused(t *testing.T) {
	if d := xsa74SessionBackendUnused(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	// No manifest backend declared -> nil.
	if d := xsa74SessionBackendUnused(fsWithFuncs()); d != nil {
		t.Errorf("no backend declared -> nil, got %v", d)
	}
	// Declared and used -> nil.
	used := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("session.Get")}})
	used.Manifest = &manifest.ProjectConfig{Session: &manifest.BuiltinBackend{Backend: "postgres"}}
	if d := xsa74SessionBackendUnused(used); d != nil {
		t.Errorf("declared+used -> nil, got %v", d)
	}
	// Declared but unused -> WARNING.
	unused := fsWithFuncs()
	unused.Manifest = &manifest.ProjectConfig{Session: &manifest.BuiltinBackend{Backend: "postgres"}}
	d := xsa74SessionBackendUnused(unused)
	if len(d) != 1 || d[0].Level != diagnostic.LevelWarning || !strings.Contains(d[0].Message, "[XSA-74]") {
		t.Fatalf("expected XSA-74 warning, got %v", d)
	}
}

func TestXsa75CacheBackendUnused(t *testing.T) {
	if d := xsa75CacheBackendUnused(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	unused := fsWithFuncs()
	unused.Manifest = &manifest.ProjectConfig{Cache: &manifest.BuiltinBackend{Backend: "redis"}}
	d := xsa75CacheBackendUnused(unused)
	if len(d) != 1 || d[0].Level != diagnostic.LevelWarning || !strings.Contains(d[0].Message, "[XSA-75]") {
		t.Fatalf("expected XSA-75 warning, got %v", d)
	}
	used := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Set")}})
	used.Manifest = &manifest.ProjectConfig{Cache: &manifest.BuiltinBackend{Backend: "redis"}}
	if d := xsa75CacheBackendUnused(used); d != nil {
		t.Errorf("declared+used -> nil, got %v", d)
	}
	if d := xsa75CacheBackendUnused(fsWithFuncs()); d != nil {
		t.Errorf("no backend -> nil, got %v", d)
	}
}

func TestXsa76FileBackendUnused(t *testing.T) {
	if d := xsa76FileBackendUnused(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	unused := fsWithFuncs()
	unused.Manifest = &manifest.ProjectConfig{File: &manifest.FileBackend{Backend: "local"}}
	d := xsa76FileBackendUnused(unused)
	if len(d) != 1 || d[0].Level != diagnostic.LevelWarning || !strings.Contains(d[0].Message, "[XSA-76]") {
		t.Fatalf("expected XSA-76 warning, got %v", d)
	}
	used := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("file.Save")}})
	used.Manifest = &manifest.ProjectConfig{File: &manifest.FileBackend{Backend: "local"}}
	if d := xsa76FileBackendUnused(used); d != nil {
		t.Errorf("declared+used -> nil, got %v", d)
	}
}

func TestXsa77QueueBackendUnused(t *testing.T) {
	if d := xsa77QueueBackendUnused(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	unused := fsWithFuncs()
	unused.Manifest = &manifest.ProjectConfig{Queue: &manifest.QueueBackend{Backend: "memory"}}
	d := xsa77QueueBackendUnused(unused)
	if len(d) != 1 || d[0].Level != diagnostic.LevelWarning || !strings.Contains(d[0].Message, "[XSA-77]") {
		t.Fatalf("expected XSA-77 warning, got %v", d)
	}
	used := fsWithFuncs(publishFunc())
	used.Manifest = &manifest.ProjectConfig{Queue: &manifest.QueueBackend{Backend: "memory"}}
	if d := xsa77QueueBackendUnused(used); d != nil {
		t.Errorf("declared+used -> nil, got %v", d)
	}
}

func TestRun(t *testing.T) {
	// Empty Fullstack -> no diagnostics (all rules short-circuit).
	if d := Run(&yongol.Fullstack{}); len(d) != 0 {
		t.Errorf("empty fs Run should yield no diags, got %v", d)
	}
	// A single XSA-75 warning aggregated through Run.
	fs := fsWithFuncs()
	fs.Manifest = &manifest.ProjectConfig{Cache: &manifest.BuiltinBackend{Backend: "redis"}}
	d := Run(fs)
	found := false
	for _, x := range d {
		if strings.Contains(x.Message, "[XSA-75]") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected XSA-75 in aggregated Run output, got %v", d)
	}
}
