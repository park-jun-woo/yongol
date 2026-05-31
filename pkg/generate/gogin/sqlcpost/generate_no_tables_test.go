//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerate — 빈 DDL/민감컬럼 유무/mkdir 에러 분기 + 파일 산출 검증
package sqlcpost

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_NoTables(t *testing.T) {
	if err := Generate(&yongol.Fullstack{}, t.TempDir()); err != nil {
		t.Errorf("empty DDLTables should return nil, got: %v", err)
	}
}
