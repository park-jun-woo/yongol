//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitOnePackage — nil/empty/inactive/unregistered/registered 분기 검증

package infra

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestEmitOnePackage(t *testing.T) {
	mctx := map[string]any{}

	t.Run("NilInterfaceSkips", func(t *testing.T) {
		fs := &yongol.Fullstack{SsacInterfaces: map[string]*ssacmeta.PackageInterface{}}
		if err := emitOnePackage(fs, "missing", mctx, "example.com/app", t.TempDir()); err != nil {
			t.Errorf("expected nil for missing iface, got: %v", err)
		}
	})

	t.Run("NoPortsSkips", func(t *testing.T) {
		fs := &yongol.Fullstack{SsacInterfaces: map[string]*ssacmeta.PackageInterface{
			"authz": {Package: "authz", Ports: nil},
		}}
		if err := emitOnePackage(fs, "authz", mctx, "example.com/app", t.TempDir()); err != nil {
			t.Errorf("expected nil for dynamic-port pkg, got: %v", err)
		}
	})

	t.Run("NoActivePortsSkips", func(t *testing.T) {
		fs := &yongol.Fullstack{SsacInterfaces: map[string]*ssacmeta.PackageInterface{
			"cache": {Package: "cache", Ports: []ssacmeta.Port{
				{Name: "CacheGet", When: "manifest.cache.backend=='off'"}, // inactive under empty mctx
			}},
		}}
		if err := emitOnePackage(fs, "cache", mctx, "example.com/app", t.TempDir()); err != nil {
			t.Errorf("expected nil when all ports inactive, got: %v", err)
		}
	})

	t.Run("UnregisteredEmitterSkips", func(t *testing.T) {
		fs := &yongol.Fullstack{SsacInterfaces: map[string]*ssacmeta.PackageInterface{
			"custompkg": {Package: "custompkg", Ports: []ssacmeta.Port{{Name: "p", When: "always"}}},
		}}
		if err := emitOnePackage(fs, "custompkg", mctx, "example.com/app", t.TempDir()); err != nil {
			t.Errorf("expected nil for unregistered emitter, got: %v", err)
		}
	})

	t.Run("RegisteredEmitterDispatched", func(t *testing.T) {
		// "cache" is registered; with an active port set lacking the required
		// CacheSet/CacheGet/CacheDelete trio, the emitter returns an error,
		// confirming dispatch reached the registered emitter.
		fs := &yongol.Fullstack{SsacInterfaces: map[string]*ssacmeta.PackageInterface{
			"cache": {Package: "cache", Ports: []ssacmeta.Port{{Name: "Bogus", When: "always"}}},
		}}
		err := emitOnePackage(fs, "cache", mctx, "example.com/app", t.TempDir())
		if err == nil {
			t.Errorf("expected error from cache emitter on incomplete ports, got nil")
		}
	})
}
