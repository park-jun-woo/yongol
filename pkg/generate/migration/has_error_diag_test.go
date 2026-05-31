//ff:func feature=migration type=test control=sequence
//ff:what TestHasErrorDiag — ERROR 레벨 진단 존재 여부
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestHasErrorDiag(t *testing.T) {
	if !hasErrorDiag([]diagnostic.Diagnostic{{Level: diagnostic.LevelWarning}, {Level: diagnostic.LevelError}}) {
		t.Error("hasErrorDiag = false, want true")
	}
	if hasErrorDiag([]diagnostic.Diagnostic{{Level: diagnostic.LevelWarning}}) {
		t.Error("hasErrorDiag = true, want false")
	}
	if hasErrorDiag(nil) {
		t.Error("hasErrorDiag(nil) = true, want false")
	}
}
