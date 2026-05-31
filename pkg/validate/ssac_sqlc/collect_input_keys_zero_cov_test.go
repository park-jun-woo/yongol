//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectInputKeys_ZeroCov(t *testing.T) {
	seq := ssac.Sequence{
		Args: []ssac.Arg{
			{Field: "OrgID"},
			{Field: ""}, // skipped
			{Field: "Page"},
		},
		Inputs: map[string]string{
			"status": "x.Status",
			"":       "ignored", // empty key skipped
		},
	}
	keys := collectInputKeys(seq)
	got := map[string]bool{}
	for _, k := range keys {
		got[k] = true
	}
	for _, want := range []string{"OrgID", "Page", "status"} {
		if !got[want] {
			t.Errorf("missing key %q in %v", want, keys)
		}
	}
	if got[""] {
		t.Error("empty key should be skipped")
	}
}
