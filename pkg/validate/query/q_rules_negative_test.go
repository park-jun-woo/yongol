//ff:func feature=validate type=test control=iteration dimension=2 topic=query-structural
//ff:what Q-* 룰 negative 테스트 — bad.sql 에서 기대한 규칙이 발화하는지 검증

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ_NegativeBadFires(t *testing.T) {
	fs := &yongol.Fullstack{SQLcQueries: loadSpecs(t, "bad.sql")}
	diags := Run(fs)
	fired := collectFiredQRules(diags, []string{"Q-02", "Q-03", "Q-04", "Q-05", "Q-06", "Q-09"})
	for _, id := range []string{"Q-02", "Q-03", "Q-05", "Q-06", "Q-09"} {
		if !fired[id] {
			t.Errorf("expected %s diagnostic, missing", id)
		}
	}
}
