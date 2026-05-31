//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what isCRUDSeq/parseDigits/splitModelMethod/planNeedsTransaction/csrfIsActive 순수 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestIsCRUDSeq(t *testing.T) {
	for _, s := range []string{ssac.SeqGet, ssac.SeqPost, ssac.SeqPut, ssac.SeqDelete} {
		if !isCRUDSeq(s) {
			t.Errorf("%q should be CRUD", s)
		}
	}
	for _, s := range []string{"auth", "call", "response", ""} {
		if isCRUDSeq(s) {
			t.Errorf("%q should not be CRUD", s)
		}
	}
}
