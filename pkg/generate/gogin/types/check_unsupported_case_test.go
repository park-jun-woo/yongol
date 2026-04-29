//ff:func feature=gen-gogin type=test-helper control=sequence
//ff:what checkUnsupportedCase — Unsupported case 1 건 검증 (Kind/Supported/sentinel marker)

package types

import (
	"strings"
	"testing"
)

// checkUnsupportedCase asserts the binding for a column the dispatcher
// must reject. The exact SqlcGoType/ApiField text is not pinned —
// presence of the "unsupported" sentinel substring suffices.
func checkUnsupportedCase(t *testing.T, c matrixCase) {
	t.Helper()
	b := MapPGType(c.col)
	if b.Kind != c.wantKind {
		t.Errorf("Kind = %v, want %v", b.Kind, c.wantKind)
	}
	if b.Supported != c.wantSupported {
		t.Errorf("Supported = %v, want %v", b.Supported, c.wantSupported)
	}
	if !strings.Contains(b.SqlcGoType, "unsupported") {
		t.Errorf("SqlcGoType = %q, expected unsupported sentinel", b.SqlcGoType)
	}
}
