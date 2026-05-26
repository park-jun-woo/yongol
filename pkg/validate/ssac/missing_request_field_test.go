//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what missingRequestField — S-60 Diagnostic 생성 검증 (File/Line/Phase/Level/Message/Advice)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestMissingRequestField(t *testing.T) {
	fn := ssac.ServiceFunc{
		FileName: "course.ssac",
	}
	seq := ssac.Sequence{
		Line: 42,
	}
	field := "instructor_id"

	got := missingRequestField(fn, seq, field)

	t.Run("File", func(t *testing.T) {
		if got.File != "course.ssac" {
			t.Errorf("File = %q, want %q", got.File, "course.ssac")
		}
	})
	t.Run("Line", func(t *testing.T) {
		if got.Line != 42 {
			t.Errorf("Line = %d, want %d", got.Line, 42)
		}
	})
	t.Run("Phase", func(t *testing.T) {
		if got.Phase != diagnostic.PhaseValidate {
			t.Errorf("Phase = %q, want %q", got.Phase, diagnostic.PhaseValidate)
		}
	})
	t.Run("Level", func(t *testing.T) {
		if got.Level != diagnostic.LevelError {
			t.Errorf("Level = %q, want %q", got.Level, diagnostic.LevelError)
		}
	})
	t.Run("Message contains S-60", func(t *testing.T) {
		if !strings.Contains(got.Message, "[S-60]") {
			t.Errorf("Message = %q, want to contain [S-60]", got.Message)
		}
	})
	t.Run("Message contains field name", func(t *testing.T) {
		if !strings.Contains(got.Message, "instructor_id") {
			t.Errorf("Message = %q, want to contain %q", got.Message, "instructor_id")
		}
	})
	t.Run("Advice non-empty", func(t *testing.T) {
		if got.Advice == "" {
			t.Error("Advice is empty, want non-empty guidance")
		}
	})
}
