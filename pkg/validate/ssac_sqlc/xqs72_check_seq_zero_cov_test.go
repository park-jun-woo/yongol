//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs72CheckSeq_ZeroCov(t *testing.T) {
	fn := ssac.ServiceFunc{}
	empty := map[string]string{}
	pm := map[string]map[string]bool{}
	dct := map[string]map[string]string{}
	qbm := map[string]sqlcparser.QuerySpec{}

	// call type → nil.
	if got := xqs72CheckSeq(fn, ssac.Sequence{Type: "call", Model: "X.Y"}, empty, pm, dct, qbm); got != nil {
		t.Error("call type should return nil")
	}
	// empty model → nil.
	if got := xqs72CheckSeq(fn, ssac.Sequence{Type: "get"}, empty, pm, dct, qbm); got != nil {
		t.Error("empty model should return nil")
	}
	// unknown query → nil.
	if got := xqs72CheckSeq(fn, ssac.Sequence{Type: "get", Model: "User.Find"}, empty, pm, dct, qbm); got != nil {
		t.Error("unknown query should return nil")
	}
}
