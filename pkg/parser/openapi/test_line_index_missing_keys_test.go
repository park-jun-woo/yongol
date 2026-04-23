//ff:func feature=openapi-parse type=test control=sequence
//ff:what LineIndex — 존재하지 않는 key lookup 은 0 반환

package openapi

import "testing"

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
