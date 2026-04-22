//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateGoReservedWords — Go 예약어 셋 정확성 회귀 방지

package ground

import "testing"

// TestPopulateGoReservedWords_Completeness guards against accidental deletion
// of reserved-word entries (downstream var-naming rules depend on this set).
func TestPopulateGoReservedWords_Completeness(t *testing.T) {
	g := newGround()
	populateGoReservedWords(g)

	set := g.Lookup["go.reserved"]
	// Spot-check core keywords (entire 25-word Go reserved list)
	for _, kw := range []string{
		"break", "case", "chan", "const", "continue",
		"default", "defer", "else", "fallthrough", "for",
		"func", "go", "goto", "if", "import",
		"interface", "map", "package", "range", "return",
		"select", "struct", "switch", "type", "var",
	} {
		if !set[kw] {
			t.Errorf("go.reserved missing %q", kw)
		}
	}
	if len(set) != 25 {
		t.Errorf("go.reserved len = %d, want 25", len(set))
	}
}
