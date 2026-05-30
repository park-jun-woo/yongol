//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestPortByName — 이름으로 port 검색 (발견/미발견) 검증

package infra

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

func TestPortByName(t *testing.T) {
	active := []ssacmeta.Port{
		{Name: "CacheSet"},
		{Name: "CacheGet"},
	}

	t.Run("Found", func(t *testing.T) {
		p := portByName(active, "CacheGet")
		if p == nil || p.Name != "CacheGet" {
			t.Errorf("expected CacheGet, got: %v", p)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		if p := portByName(active, "CacheDelete"); p != nil {
			t.Errorf("expected nil for missing port, got: %v", p)
		}
	})

	t.Run("EmptyList", func(t *testing.T) {
		if p := portByName(nil, "CacheSet"); p != nil {
			t.Errorf("expected nil for empty list, got: %v", p)
		}
	})
}
