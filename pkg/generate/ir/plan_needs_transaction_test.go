//ff:func feature=gen-ir type=test control=sequence
//ff:what isCRUDSeq/parseDigits/splitModelMethod/planNeedsTransaction/csrfIsActive 순수 헬퍼
package ir

import (
	"testing"
)

func TestPlanNeedsTransaction(t *testing.T) {
	if !planNeedsTransaction([]Op{{Kind: OpGet}, {Kind: OpPost}}) {
		t.Errorf("post should need transaction")
	}
	if !planNeedsTransaction([]Op{{Kind: OpDelete}}) {
		t.Errorf("delete should need transaction")
	}
	if planNeedsTransaction([]Op{{Kind: OpGet}}) {
		t.Errorf("get-only should not need transaction")
	}
	if planNeedsTransaction(nil) {
		t.Errorf("empty should not need transaction")
	}
}
