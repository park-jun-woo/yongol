//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestScanLine — 라인 → (제어종류, brace delta) 환산 검증

package ffannot

import "testing"

func TestScanLine(t *testing.T) {
	t.Run("EmptyLine", func(t *testing.T) {
		kind, delta := scanLine("   ", 0)
		if kind != "" || delta != 0 {
			t.Errorf("empty line: got (%q,%d), want (\"\",0)", kind, delta)
		}
	})

	t.Run("ForAtDepth0", func(t *testing.T) {
		kind, delta := scanLine("for i := range xs {", 0)
		if kind != ControlIteration {
			t.Errorf("expected iteration, got %q", kind)
		}
		if delta != 1 {
			t.Errorf("expected delta 1, got %d", delta)
		}
	})

	t.Run("SwitchAtDepth0", func(t *testing.T) {
		kind, _ := scanLine("switch x {", 0)
		if kind != ControlSelection {
			t.Errorf("expected selection, got %q", kind)
		}
	})

	t.Run("ForNotClassifiedWhenNested", func(t *testing.T) {
		// depth > 0 -> classifyDepth0 not consulted; kind stays "".
		kind, delta := scanLine("for i := range xs {", 1)
		if kind != "" {
			t.Errorf("expected no kind at depth>0, got %q", kind)
		}
		if delta != 1 {
			t.Errorf("expected delta 1, got %d", delta)
		}
	})
}
