//ff:func feature=manifest type=test control=sequence
//ff:what LineIndex nil receiver + missing-key 경로가 0 을 반환하는지 검증

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

func TestLineIndex_MissingKeys(t *testing.T) {
	l := &LineIndex{
		Paths:            map[string]int{},
		Schemas:          map[string]int{},
		RequestFields:    map[string]map[string]int{},
		ResponseFields:   map[string]map[string]int{},
		SchemaProperties: map[string]map[string]int{},
		Operations:       map[string]int{},
	}
	if got := l.PathLine("/nope"); got != 0 {
		t.Errorf("missing PathLine = %d, want 0", got)
	}
	if got := l.SchemaLine("Nope"); got != 0 {
		t.Errorf("missing SchemaLine = %d, want 0", got)
	}
	if got := l.RequestFieldLine("Nope", "f"); got != 0 {
		t.Errorf("missing op RequestFieldLine = %d, want 0", got)
	}
	if got := l.ResponseFieldLine("Nope", "f"); got != 0 {
		t.Errorf("missing op ResponseFieldLine = %d, want 0", got)
	}
	if got := l.SchemaPropertyLine("Nope", "f"); got != 0 {
		t.Errorf("missing schema SchemaPropertyLine = %d, want 0", got)
	}
	if got := l.OperationLine("Nope"); got != 0 {
		t.Errorf("missing OperationLine = %d, want 0", got)
	}
}

func TestLineIndex_PathLine_Populated(t *testing.T) {
	path := writeFixture(t)
	idx, err := BuildLineIndex(path)
	if err != nil {
		t.Fatalf("BuildLineIndex: %v", err)
	}
	if got := idx.PathLine("/login"); got != 14 {
		t.Errorf("PathLine(/login) = %d, want 14", got)
	}
	if got := idx.PathLine("/missing"); got != 0 {
		t.Errorf("PathLine(/missing) = %d, want 0", got)
	}
}
