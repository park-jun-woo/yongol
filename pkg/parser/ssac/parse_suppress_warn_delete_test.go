//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @delete! SuppressWarn 파싱 검증

package ssac

import "testing"

func TestParseSuppressWarnDelete(t *testing.T) {
	src := `package service

// @delete! Room.DeleteAll()
func DeleteAll() {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqDelete)
	assertEqual(t, "Model", seq.Model, "Room.DeleteAll")
	if !seq.SuppressWarn {
		t.Error("expected SuppressWarn=true")
	}
}
