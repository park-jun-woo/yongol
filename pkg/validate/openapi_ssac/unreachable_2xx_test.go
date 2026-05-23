//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what unreachable2xx — 선택됨 제외/미선택 포함 검증

package openapi_ssac

import "testing"

func TestUnreachable2xx(t *testing.T) {
	t.Run("empty declared returns empty", func(t *testing.T) {
		got := unreachable2xx(map[int]bool{}, 200)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("selected code excluded", func(t *testing.T) {
		declared := map[int]bool{200: true, 201: true, 204: true}
		got := unreachable2xx(declared, 200)
		if len(got) != 2 {
			t.Fatalf("expected 2, got %d: %v", len(got), got)
		}
		if got[200] {
			t.Error("200 should be excluded")
		}
		if !got[201] || !got[204] {
			t.Errorf("expected 201 and 204: %v", got)
		}
	})

	t.Run("selected not in declared returns all", func(t *testing.T) {
		declared := map[int]bool{200: true}
		got := unreachable2xx(declared, 201)
		if len(got) != 1 || !got[200] {
			t.Errorf("expected {200: true}, got %v", got)
		}
	})
}
