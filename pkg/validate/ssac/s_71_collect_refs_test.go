//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-71 — s71CollectRefs 단위 테스트 (시퀀스에서 변수 참조 수집)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestS71CollectRefs(t *testing.T) {
	t.Run("Collects_inputs", func(t *testing.T) {
		seq := ssac.Sequence{Inputs: map[string]string{"a": "x.ID", "b": "y.Name"}}
		refs := s71CollectRefs(seq)
		if len(refs) != 2 {
			t.Errorf("expected 2 refs, got %d", len(refs))
		}
	})
	t.Run("Collects_target", func(t *testing.T) {
		seq := ssac.Sequence{Target: "course"}
		refs := s71CollectRefs(seq)
		if len(refs) != 1 || refs[0] != "course" {
			t.Errorf("expected [course], got %v", refs)
		}
	})
	t.Run("Collects_email_expr", func(t *testing.T) {
		seq := ssac.Sequence{EmailExpr: "request.Email"}
		refs := s71CollectRefs(seq)
		if len(refs) != 1 || refs[0] != "request.Email" {
			t.Errorf("expected [request.Email], got %v", refs)
		}
	})
	t.Run("Collects_password_expr", func(t *testing.T) {
		seq := ssac.Sequence{PasswordExpr: "request.Password"}
		refs := s71CollectRefs(seq)
		if len(refs) != 1 || refs[0] != "request.Password" {
			t.Errorf("expected [request.Password], got %v", refs)
		}
	})
	t.Run("Collects_fields", func(t *testing.T) {
		seq := ssac.Sequence{Fields: map[string]string{"x": "a", "y": "b"}}
		refs := s71CollectRefs(seq)
		if len(refs) != 2 {
			t.Errorf("expected 2 refs, got %d", len(refs))
		}
	})
	t.Run("Empty_seq_returns_nil", func(t *testing.T) {
		refs := s71CollectRefs(ssac.Sequence{})
		if len(refs) != 0 {
			t.Errorf("expected 0 refs, got %d", len(refs))
		}
	})
}
