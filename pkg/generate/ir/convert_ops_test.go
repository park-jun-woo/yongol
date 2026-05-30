//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what convertDelete/convertPut/convertEmpty/convertExists/matchFollowingGuard/resolveVar/convertInputsToFieldArgs

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertDelete(t *testing.T) {
	op := convertDelete(ssac.Sequence{Model: "Reservation.Delete", Inputs: map[string]string{"ID": "request.ID"}})
	if op.Kind != OpDelete || op.Delete == nil {
		t.Fatalf("op = %+v", op)
	}
	if op.Delete.Model != "Reservation" || op.Delete.Method != "Delete" {
		t.Errorf("model/method = %q/%q", op.Delete.Model, op.Delete.Method)
	}
	if len(op.Delete.Args) != 1 {
		t.Errorf("args = %+v", op.Delete.Args)
	}
}

func TestConvertPut(t *testing.T) {
	op := convertPut(ssac.Sequence{Model: "Course.Update"})
	if op.Kind != OpPut || op.Put == nil || op.Put.Model != "Course" || op.Put.Method != "Update" {
		t.Errorf("op = %+v", op)
	}
}

func TestConvertEmpty(t *testing.T) {
	// explicit status
	op := convertEmpty(ssac.Sequence{Target: "course", Message: "not found", ErrStatus: 410})
	if op.Kind != OpEmpty || op.Empty.VarName != "course" || op.Empty.Message != "not found" || op.Empty.StatusCode != 410 {
		t.Errorf("op = %+v", op.Empty)
	}
	// default status 404
	def := convertEmpty(ssac.Sequence{Target: "x"})
	if def.Empty.StatusCode != 404 {
		t.Errorf("default status = %d, want 404", def.Empty.StatusCode)
	}
}

func TestConvertExists(t *testing.T) {
	op := convertExists(ssac.Sequence{Target: "dup", Message: "conflict"})
	if op.Kind != OpExists || op.Exists.VarName != "dup" || op.Exists.StatusCode != 409 {
		t.Errorf("op = %+v", op.Exists)
	}
	withStatus := convertExists(ssac.Sequence{Target: "y", ErrStatus: 422})
	if withStatus.Exists.StatusCode != 422 {
		t.Errorf("status = %d, want 422", withStatus.Exists.StatusCode)
	}
}

func TestMatchFollowingGuard(t *testing.T) {
	empty := Op{Kind: OpEmpty, Empty: &EmptyOp{VarName: "course"}}
	if got := matchFollowingGuard(empty, "course"); got != OpEmpty {
		t.Errorf("empty match = %v, want OpEmpty", got)
	}
	exists := Op{Kind: OpExists, Exists: &ExistsOp{VarName: "dup"}}
	if got := matchFollowingGuard(exists, "dup"); got != OpExists {
		t.Errorf("exists match = %v, want OpExists", got)
	}
	// non-matching var name -> OpGet (zero)
	if got := matchFollowingGuard(empty, "other"); got != OpGet {
		t.Errorf("non-match = %v, want OpGet", got)
	}
	// non-guard op -> OpGet
	if got := matchFollowingGuard(Op{Kind: OpPost}, "x"); got != OpGet {
		t.Errorf("non-guard = %v, want OpGet", got)
	}
}

func TestResolveVar(t *testing.T) {
	declared := map[string]bool{}
	if got := resolveVar("user", declared); got != "user" {
		t.Errorf("first use = %q, want user", got)
	}
	// second use collides -> _result suffix
	if got := resolveVar("user", declared); got != "user_result" {
		t.Errorf("second use = %q, want user_result", got)
	}
	// third use collides again
	if got := resolveVar("user", declared); got != "user_result_result" {
		t.Errorf("third use = %q, want user_result_result", got)
	}
}

func TestConvertInputsToFieldArgs(t *testing.T) {
	if got := convertInputsToFieldArgs(nil); got != nil {
		t.Errorf("nil inputs = %+v, want nil", got)
	}
	args := convertInputsToFieldArgs(map[string]string{"Zeta": "request.Z", "Alpha": "request.A"})
	if len(args) != 2 {
		t.Fatalf("len = %d, want 2", len(args))
	}
	// keys sorted deterministically: Alpha before Zeta
	if args[0].Key != "Alpha" || args[1].Key != "Zeta" {
		t.Errorf("keys not sorted: %q, %q", args[0].Key, args[1].Key)
	}
}
