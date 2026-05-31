//ff:func feature=gen-gogin type=test control=sequence
//ff:what collectSensitiveKeys — DDL `-- @sensitive` 컬럼명을 sorted 리스트로 수집
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCollectSensitiveKeys_None(t *testing.T) {
	got := collectSensitiveKeys(&yongol.Fullstack{})
	if len(got) != 0 {
		t.Errorf("expected empty, got %v", got)
	}
}
