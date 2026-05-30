//ff:func feature=validate type=test control=selection topic=ssac-statemachine
//ff:what TestMakeStateTypeDiag — makeStateTypeDiag XSM-71 진단 생성 분기 검증

package ssac_statemachine

import (
	"strings"
	"testing"
)

func TestMakeStateTypeDiag(t *testing.T) {
	t.Run("empty source type returns nil", func(t *testing.T) {
		if d := makeStateTypeDiag("f.ssac", 3, "status", ""); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("string-compatible returns nil", func(t *testing.T) {
		if d := makeStateTypeDiag("f.ssac", 3, "status", "string"); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("incompatible returns diagnostic", func(t *testing.T) {
		d := makeStateTypeDiag("f.ssac", 7, "status", "uuid.UUID")
		if d == nil {
			t.Fatal("expected diagnostic, got nil")
		}
		if d.File != "f.ssac" || d.Line != 7 {
			t.Errorf("unexpected location: %+v", d)
		}
		if !strings.Contains(d.Message, "XSM-71") || !strings.Contains(d.Message, "uuid.UUID") {
			t.Errorf("unexpected message: %q", d.Message)
		}
	})
}
