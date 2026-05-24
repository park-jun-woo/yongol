//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what hasFKRef — FK 참조 검출 (args/inputs/implicit/same model/no ref) 검증

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestHasFKRef(t *testing.T) {
	declared := map[string]bool{"user": true, "order": true}
	types := map[string]string{"user": "User", "order": "Order"}

	t.Run("arg references foreign model", func(t *testing.T) {
		seq := parsessac.Sequence{
			Args: []parsessac.Arg{{Source: "user", Field: "ID"}},
		}
		if !hasFKRef(seq, declared, types, "Order") {
			t.Error("expected true for FK ref via args")
		}
	})

	t.Run("arg references same model", func(t *testing.T) {
		seq := parsessac.Sequence{
			Args: []parsessac.Arg{{Source: "user", Field: "ID"}},
		}
		if hasFKRef(seq, declared, types, "User") {
			t.Error("expected false for same model")
		}
	})

	t.Run("implicit source skipped", func(t *testing.T) {
		seq := parsessac.Sequence{
			Args: []parsessac.Arg{{Source: "request", Field: "ID"}},
		}
		if hasFKRef(seq, declared, types, "Order") {
			t.Error("expected false for implicit source")
		}
	})

	t.Run("input references foreign model", func(t *testing.T) {
		seq := parsessac.Sequence{
			Inputs: map[string]string{"user_id": "user.ID"},
		}
		if !hasFKRef(seq, declared, types, "Order") {
			t.Error("expected true for FK ref via inputs")
		}
	})

	t.Run("no args or inputs", func(t *testing.T) {
		seq := parsessac.Sequence{}
		if hasFKRef(seq, declared, types, "Order") {
			t.Error("expected false")
		}
	})

	t.Run("undeclared var in args", func(t *testing.T) {
		seq := parsessac.Sequence{
			Args: []parsessac.Arg{{Source: "unknown"}},
		}
		if hasFKRef(seq, declared, types, "Order") {
			t.Error("expected false for undeclared var")
		}
	})

	t.Run("implicit source in inputs skipped", func(t *testing.T) {
		seq := parsessac.Sequence{
			Inputs: map[string]string{"id": "request.ID"},
		}
		if hasFKRef(seq, declared, types, "Order") {
			t.Error("expected false for implicit source in inputs")
		}
	})

	t.Run("literal in inputs skipped", func(t *testing.T) {
		seq := parsessac.Sequence{
			Inputs: map[string]string{"status": `"active"`},
		}
		if hasFKRef(seq, declared, types, "Order") {
			t.Error("expected false for literal in inputs")
		}
	})

	t.Run("empty source in args skipped", func(t *testing.T) {
		seq := parsessac.Sequence{
			Args: []parsessac.Arg{{Source: ""}},
		}
		if hasFKRef(seq, declared, types, "Order") {
			t.Error("expected false for empty source")
		}
	})
}
