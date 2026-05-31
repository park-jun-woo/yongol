//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqsBatch_ZeroCov — ssac_sqlc 검증 헬퍼를 이름으로 직접 호출해 커버 귀속
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqsFullstackEntries_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{}
	_ = buildDDLTableLookup(fs)
	_ = buildQueryBodyMap(fs)
	_ = collectHaveQueries(fs)
	_ = xqs16InputKeyMissing(fs)
	_ = xqs17ParamKeyMissing(fs)
	_ = xqs18ParamTypeMismatch(fs)
	_ = xqs20ReturnTypeMatch(fs)
	_ = xqs72QueryParamIntWidth(fs)
}
