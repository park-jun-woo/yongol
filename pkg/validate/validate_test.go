//ff:func feature=validate type=test control=sequence
//ff:what TestValidate — ground 자동 빌드 / WithGround 주입 / WithArtsDir 분기 검증

package validate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestValidate(t *testing.T) {
	// Default: no options → ground built internally, no artsDir step.
	fs := &yongol.Fullstack{}
	r := Validate(fs)
	if r == nil {
		t.Fatal("expected non-nil report")
	}
	baseSteps := len(r.Steps)
	if baseSteps == 0 {
		t.Fatal("expected steps from allSteps()")
	}

	// WithGround injected (skip internal Build) + WithArtsDir set.
	fs2 := &yongol.Fullstack{}
	g := ground.Build(fs2)
	r2 := Validate(fs2, WithGround(g), WithArtsDir(t.TempDir()))
	// artsDir non-empty → one extra "contract" step appended.
	if len(r2.Steps) != baseSteps+1 {
		t.Fatalf("expected %d steps with artsDir, got %d", baseSteps+1, len(r2.Steps))
	}
	last := r2.Steps[len(r2.Steps)-1]
	if last.Name != "contract" {
		t.Fatalf("expected final step named contract, got %q", last.Name)
	}
}
