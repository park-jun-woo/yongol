//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildQueryParamMap_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "UserFindByEmail", Params: []string{"Email"}},
			{Name: "ListAll", Params: nil}, // skipped (no params)
		},
	}
	m := buildQueryParamMap(fs)
	if _, ok := m["ListAll"]; ok {
		t.Error("query with no params should be skipped")
	}
	set := m["UserFindByEmail"]
	if !set["Email"] {
		t.Errorf("expected Email param, got %v", set)
	}
}
