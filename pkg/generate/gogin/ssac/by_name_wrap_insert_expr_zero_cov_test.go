//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac

import (
	"testing"
)

func TestByNameWrapInsertExpr_ZeroCov(t *testing.T) {
	g := &methodGen{}
	// alreadyPgtype → passthrough.
	if got, imp := g.wrapInsertExpr("k", "row.Field", true, "request.field"); got != "row.Field" || imp != nil {
		t.Errorf("alreadyPgtype = %q %v", got, imp)
	}
	// no SQLc column → passthrough (lookupSQLCMethodColumn nil).
	if got, _ := g.wrapInsertExpr("unknown", "v", false, "v"); got != "v" {
		t.Errorf("no-col passthrough = %q", got)
	}
}
