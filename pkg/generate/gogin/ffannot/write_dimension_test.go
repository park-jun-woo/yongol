//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWriteDimension — iteration 일 때만 dimension 덧붙임 + 기본값 1 검증

package ffannot

import (
	"strings"
	"testing"
)

func TestWriteDimension(t *testing.T) {
	t.Run("NonIterationNoOp", func(t *testing.T) {
		var sb strings.Builder
		writeDimension(&sb, FuncAnnot{Control: "sequence", Dimension: 3})
		if sb.String() != "" {
			t.Errorf("expected no output for non-iteration, got %q", sb.String())
		}
	})

	t.Run("IterationExplicitDim", func(t *testing.T) {
		var sb strings.Builder
		writeDimension(&sb, FuncAnnot{Control: "iteration", Dimension: 2})
		if sb.String() != " dimension=2" {
			t.Errorf("expected ' dimension=2', got %q", sb.String())
		}
	})

	t.Run("IterationDefaultsToOne", func(t *testing.T) {
		var sb strings.Builder
		writeDimension(&sb, FuncAnnot{Control: "iteration", Dimension: 0})
		if sb.String() != " dimension=1" {
			t.Errorf("expected ' dimension=1', got %q", sb.String())
		}
	})
}
