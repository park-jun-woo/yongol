//ff:func feature=stml-parse type=parser control=sequence
//ff:what assertField — test helper to verify FieldBind name, tag, and type

package stml

import (
	"testing"
)

func assertField(t *testing.T, f FieldBind, name, tag, typ string) {
	t.Helper()
	if f.Name != name {
		t.Errorf("Field.Name = %q, want %q", f.Name, name)
	}
	if f.Tag != tag {
		t.Errorf("Field.Tag = %q, want %q", f.Tag, tag)
	}
	if f.Type != typ {
		t.Errorf("Field.Type = %q, want %q", f.Type, typ)
	}
}
