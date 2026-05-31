//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestConvertSequence_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{}
	cases := []struct {
		seqType string
		seq     ssac.Sequence
		want    OpKind
	}{
		{ssac.SeqGet, ssac.Sequence{Type: ssac.SeqGet, Model: "A.B"}, OpGet},
		{ssac.SeqPost, ssac.Sequence{Type: ssac.SeqPost, Model: "A.B"}, OpPost},
		{ssac.SeqPut, ssac.Sequence{Type: ssac.SeqPut, Model: "A.B"}, OpPut},
		{ssac.SeqDelete, ssac.Sequence{Type: ssac.SeqDelete, Model: "A.B"}, OpDelete},
		{ssac.SeqEmpty, ssac.Sequence{Type: ssac.SeqEmpty, Target: "x"}, OpEmpty},
		{ssac.SeqExists, ssac.Sequence{Type: ssac.SeqExists, Target: "x"}, OpExists},
		{ssac.SeqAuth, ssac.Sequence{Type: ssac.SeqAuth, Resource: "p"}, OpAuth},
		{ssac.SeqState, ssac.Sequence{Type: ssac.SeqState, DiagramID: "d", Transition: "t"}, OpState},
		{ssac.SeqCall, ssac.Sequence{Type: ssac.SeqCall, Model: "A.B"}, OpCall},
		{ssac.SeqEval, ssac.Sequence{Type: ssac.SeqEval, Model: "A.B"}, OpEval},
		{ssac.SeqPublish, ssac.Sequence{Type: ssac.SeqPublish, Topic: "t"}, OpPublish},
		{ssac.SeqVerifyPassword, ssac.Sequence{Type: ssac.SeqVerifyPassword, Model: "User"}, OpVerifyPassword},
		{ssac.SeqResponse, ssac.Sequence{Type: ssac.SeqResponse, Target: "x"}, OpResponse},
	}
	for _, c := range cases {
		op, err := convertSequence(c.seq, fs)
		if err != nil {
			t.Errorf("%s: unexpected error %v", c.seqType, err)
			continue
		}
		if op.Kind != c.want {
			t.Errorf("%s: Kind = %d, want %d", c.seqType, op.Kind, c.want)
		}
	}
	// unknown type → error
	if _, err := convertSequence(ssac.Sequence{Type: "bogus"}, fs); err == nil {
		t.Error("expected error for unknown sequence type")
	}
}
