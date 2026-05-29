//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXnc90_MemoryBackend_NoDiagnostic — memory backend 는 XNC-90 을 트리거하지 않음

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnc90_MemoryBackend_NoDiagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Cache: &pmanifest.BuiltinBackend{Backend: "memory"},
		},
	}
	if diags := xnc90CacheBackendRequiresSQLC(fs); len(diags) != 0 {
		t.Errorf("memory backend must not trigger XNC-90: %+v", diags)
	}
}
