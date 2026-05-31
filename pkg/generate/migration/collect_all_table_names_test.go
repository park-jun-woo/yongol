//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 파이프라인 함수별 named 테스트 — tsma 함수명 매칭용 (parse/diff/emit/tokenizer 커버)
package migration

import (
	"testing"
)

func TestCollectAllTableNames(t *testing.T) {
	prev, _, _ := migPipelineSchemas(t)
	curr := NewSchema()
	if len(collectAllTableNames(prev, curr)) == 0 {
		t.Errorf("expected table names")
	}
}
