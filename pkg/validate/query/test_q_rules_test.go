//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what Q-* 룰 golden 테스트 — bad 쿼리 없이 실행 시 Q-* 진단 미발화 확인

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ_GoldenAllPass(t *testing.T) {
	fs := &yongol.Fullstack{SQLcQueries: loadSpecs(t, "good.sql")}
	diags := Run(fs)
	for _, d := range diags {
		if strings.Contains(d.Message, "[Q-") {
			t.Errorf("unexpected Q-* diag: %s", d.Message)
		}
	}
}
