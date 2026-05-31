//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs20EligibleSeqType(t *testing.T) {
	for _, ty := range []string{ssacparser.SeqGet, ssacparser.SeqPost, ssacparser.SeqPut} {
		if !xqs20EligibleSeqType(ty) {
			t.Errorf("%q should be eligible", ty)
		}
	}
	if xqs20EligibleSeqType(ssacparser.SeqDelete) {
		t.Error("delete should not be eligible")
	}
}
