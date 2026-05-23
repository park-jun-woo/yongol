//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what canonicalAdvice — ssac interface.yaml canonical DDL/queries 렌더 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCanonicalAdvice(t *testing.T) {
	t.Run("nil_fs_fallback", func(t *testing.T) {
		got := canonicalAdvice(nil, "cache")
		if !strings.Contains(got, "interface.yaml") {
			t.Errorf("expected fallback text, got %q", got)
		}
	})

	t.Run("nil_interfaces_fallback", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		got := canonicalAdvice(fs, "cache")
		if !strings.Contains(got, "interface.yaml") {
			t.Errorf("expected fallback text, got %q", got)
		}
	})

	t.Run("pkg_not_found_fallback", func(t *testing.T) {
		fs := &yongol.Fullstack{SsacInterfaces: map[string]*ssacmeta.PackageInterface{}}
		got := canonicalAdvice(fs, "cache")
		if !strings.Contains(got, "interface.yaml") {
			t.Errorf("expected fallback text, got %q", got)
		}
	})

	t.Run("renders_ddl_and_queries", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SsacInterfaces: map[string]*ssacmeta.PackageInterface{
				"cache": {
					CanonicalDDL:     "CREATE TABLE cache_entries (...);",
					CanonicalQueries: "-- name: GetCacheEntry\nSELECT * FROM cache_entries;",
				},
			},
		}
		got := canonicalAdvice(fs, "cache")
		if !strings.Contains(got, "cache.sql") {
			t.Errorf("expected DDL file reference, got %q", got)
		}
		if !strings.Contains(got, "CREATE TABLE") {
			t.Errorf("expected DDL content, got %q", got)
		}
		if !strings.Contains(got, "GetCacheEntry") {
			t.Errorf("expected queries content, got %q", got)
		}
	})

	t.Run("ddl_only", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SsacInterfaces: map[string]*ssacmeta.PackageInterface{
				"session": {CanonicalDDL: "CREATE TABLE sessions (...);"},
			},
		}
		got := canonicalAdvice(fs, "session")
		if !strings.Contains(got, "session.sql") {
			t.Errorf("expected DDL reference, got %q", got)
		}
		if strings.Contains(got, "queries/session.sql") {
			t.Errorf("should not contain queries when empty")
		}
	})
}
