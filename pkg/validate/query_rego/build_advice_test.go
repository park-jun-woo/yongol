//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what buildAdvice — 직접/JOIN 형태 advice 생성 검증

package query_rego

import (
	"strings"
	"testing"
)

func TestBuildAdvice(t *testing.T) {
	t.Run("direct mapping produces single-table SELECT", func(t *testing.T) {
		advice := buildAdvice("GetOwner", "orders", "user_id", "", "")
		if !strings.Contains(advice, "SELECT user_id FROM orders WHERE id = @id") {
			t.Errorf("expected single-table SELECT, got %s", advice)
		}
		if !strings.Contains(advice, "-- name: GetOwner :one") {
			t.Errorf("expected name annotation, got %s", advice)
		}
		if !strings.Contains(advice, "orders.sql") {
			t.Errorf("expected file reference, got %s", advice)
		}
	})

	t.Run("via annotation produces JOIN form", func(t *testing.T) {
		advice := buildAdvice("GetOwner", "comments", "user_id", "posts", "comment_id")
		if !strings.Contains(advice, "JOIN posts") {
			t.Errorf("expected JOIN, got %s", advice)
		}
		if !strings.Contains(advice, "c.user_id") {
			t.Errorf("expected c.user_id, got %s", advice)
		}
		if !strings.Contains(advice, "l.comment_id = c.id") {
			t.Errorf("expected join condition, got %s", advice)
		}
		if !strings.Contains(advice, "-- name: GetOwner :one") {
			t.Errorf("expected name annotation, got %s", advice)
		}
	})

	t.Run("join with empty joinFK falls back to direct", func(t *testing.T) {
		advice := buildAdvice("GetOwner", "orders", "user_id", "items", "")
		if strings.Contains(advice, "JOIN") {
			t.Errorf("expected no JOIN when joinFK is empty, got %s", advice)
		}
	})
}
