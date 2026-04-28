//ff:func feature=validate type=test-helper control=sequence topic=ssac-sqlc
//ff:what cacheInterface — 테스트용 ssacmeta.PackageInterface fixture (cache)

package ssac_sqlc

import "github.com/park-jun-woo/yongol/pkg/ssacmeta"

// cacheInterface returns a minimal ssacmeta.PackageInterface fixture for the
// ssac cache package. Mirrors the real ssac/pkg/cache/interface.yaml so the
// unit test stays faithful to the authoritative catalog.
func cacheInterface() *ssacmeta.PackageInterface {
	return &ssacmeta.PackageInterface{
		Package: "cache",
		Ports: []ssacmeta.Port{
			{Name: "CacheSet", UsedBy: []string{"Set"}},
			{Name: "CacheGet", UsedBy: []string{"Get"}},
			{Name: "CacheDelete", UsedBy: []string{"Delete"}},
		},
	}
}
