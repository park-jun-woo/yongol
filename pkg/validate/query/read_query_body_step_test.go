//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what readQueryBodyStep — 한 라인 처리 (시작 라인/본문/종료/escape/비쿼리) 검증

package query

import "testing"

func TestReadQueryBodyStep(t *testing.T) {
	t.Run("start line sets header", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		inQuery := false
		done := readQueryBodyStep(body, "-- name: GetUser :one", 5, 5, &inQuery)
		if done {
			t.Error("expected false")
		}
		if !inQuery {
			t.Error("expected inQuery to be true")
		}
		if body.Header != "-- name: GetUser :one" {
			t.Errorf("expected header, got %q", body.Header)
		}
	})

	t.Run("before start line does nothing", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		inQuery := false
		done := readQueryBodyStep(body, "some line", 3, 5, &inQuery)
		if done {
			t.Error("expected false")
		}
		if inQuery {
			t.Error("expected inQuery to remain false")
		}
	})

	t.Run("body line is appended", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		inQuery := true
		done := readQueryBodyStep(body, "SELECT * FROM users;", 6, 5, &inQuery)
		if done {
			t.Error("expected false")
		}
		if len(body.Lines) != 1 {
			t.Fatalf("expected 1 line, got %d", len(body.Lines))
		}
		if body.Lines[0] != "SELECT * FROM users;" {
			t.Errorf("expected SQL line, got %q", body.Lines[0])
		}
	})

	t.Run("next name annotation stops", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		inQuery := true
		done := readQueryBodyStep(body, "-- name: ListUsers :many", 7, 5, &inQuery)
		if !done {
			t.Error("expected true (stop)")
		}
	})

	t.Run("escape comment is registered", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		inQuery := true
		done := readQueryBodyStep(body, "-- +no-pagination", 6, 5, &inQuery)
		if done {
			t.Error("expected false")
		}
		if !body.Escapes["@no-pagination"] {
			t.Error("expected @no-pagination escape")
		}
	})

	t.Run("regular comment is appended", func(t *testing.T) {
		body := &queryBody{Escapes: make(map[string]bool)}
		inQuery := true
		done := readQueryBodyStep(body, "-- just a comment", 6, 5, &inQuery)
		if done {
			t.Error("expected false")
		}
		if len(body.Lines) != 1 {
			t.Fatalf("expected 1 line, got %d", len(body.Lines))
		}
	})
}
