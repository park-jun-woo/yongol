//ff:func feature=manifest type=parser control=sequence
//ff:what LineIndex 메서드들이 nil receiver 에서도 panic 없이 0 을 반환하는지 검증

package openapi

import "testing"

func TestBuildLineIndex_NilSafe(t *testing.T) {
	var l *LineIndex
	if got := l.OperationLine("X"); got != 0 {
		t.Errorf("nil OperationLine = %d, want 0", got)
	}
	if got := l.SchemaPropertyLine("X", "y"); got != 0 {
		t.Errorf("nil SchemaPropertyLine = %d, want 0", got)
	}
}
