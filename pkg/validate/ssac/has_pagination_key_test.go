//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what hasPaginationKey — pagination 키 존재 여부 검증 (found/not found/empty)

package ssac

import "testing"

func TestHasPaginationKey(t *testing.T) {
	t.Run("empty inputs returns false", func(t *testing.T) {
		if hasPaginationKey(nil) {
			t.Error("expected false")
		}
	})

	t.Run("Page found", func(t *testing.T) {
		if !hasPaginationKey(map[string]string{"Page": "query.Page"}) {
			t.Error("expected true")
		}
	})

	t.Run("PerPage found", func(t *testing.T) {
		if !hasPaginationKey(map[string]string{"PerPage": "query.PerPage"}) {
			t.Error("expected true")
		}
	})

	t.Run("Cursor found", func(t *testing.T) {
		if !hasPaginationKey(map[string]string{"Cursor": "query.Cursor"}) {
			t.Error("expected true")
		}
	})

	t.Run("no pagination key", func(t *testing.T) {
		if hasPaginationKey(map[string]string{"status": "active"}) {
			t.Error("expected false")
		}
	})
}
