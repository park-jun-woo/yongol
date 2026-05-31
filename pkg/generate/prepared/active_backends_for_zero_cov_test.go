//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

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
