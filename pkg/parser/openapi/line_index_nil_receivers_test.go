//ff:func feature=openapi-parse type=test control=sequence
//ff:what LineIndex — nil receiver 에서 모든 lookup 이 0 반환

package openapi

import "testing"

func TestLineIndex_NilReceivers(t *testing.T) {
	var l *LineIndex
	if got := l.PathLine("/x"); got != 0 {
		t.Errorf("nil PathLine = %d", got)
	}
	if got := l.SchemaLine("X"); got != 0 {
		t.Errorf("nil SchemaLine = %d", got)
	}
	if got := l.RequestFieldLine("Op", "f"); got != 0 {
		t.Errorf("nil RequestFieldLine = %d", got)
	}
	if got := l.ResponseFieldLine("Op", "f"); got != 0 {
		t.Errorf("nil ResponseFieldLine = %d", got)
	}
}
