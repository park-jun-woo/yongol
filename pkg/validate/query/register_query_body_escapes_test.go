//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what registerQueryBodyEscapes — escape 마커 등록 (+form/@form/미매칭/복수) 검증

package query

import "testing"

func TestRegisterQueryBodyEscapes(t *testing.T) {
	t.Run("plus form registers", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		registerQueryBodyEscapes(body, "-- +no-pagination")
		if !body.Escapes["@no-pagination"] {
			t.Error("expected @no-pagination")
		}
		if !body.HasStop {
			t.Error("expected HasStop true")
		}
	})

	t.Run("at form registers", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		registerQueryBodyEscapes(body, "-- @allow-truncate")
		if !body.Escapes["@allow-truncate"] {
			t.Error("expected @allow-truncate")
		}
	})

	t.Run("allow-sensitive registers", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		registerQueryBodyEscapes(body, "-- +allow-sensitive")
		if !body.Escapes["@allow-sensitive"] {
			t.Error("expected @allow-sensitive")
		}
	})

	t.Run("skip-column-check registers", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		registerQueryBodyEscapes(body, "-- @skip-column-check")
		if !body.Escapes["@skip-column-check"] {
			t.Error("expected @skip-column-check")
		}
	})

	t.Run("unrelated comment does not register", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		registerQueryBodyEscapes(body, "-- just a regular comment")
		if body.HasStop {
			t.Error("expected HasStop false")
		}
		if len(body.Escapes) != 0 {
			t.Errorf("expected empty escapes, got %v", body.Escapes)
		}
	})

	t.Run("multiple markers in one line", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		registerQueryBodyEscapes(body, "-- +no-pagination +allow-truncate")
		if !body.Escapes["@no-pagination"] {
			t.Error("expected @no-pagination")
		}
		if !body.Escapes["@allow-truncate"] {
			t.Error("expected @allow-truncate")
		}
	})
}
