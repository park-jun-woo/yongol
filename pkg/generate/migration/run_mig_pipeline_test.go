//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 파이프라인 함수별 named 테스트 — tsma 함수명 매칭용 (parse/diff/emit/tokenizer 커버)
package migration

import (
	"testing"
)

func runMigPipeline(t *testing.T) []Operation {
	t.Helper()
	prev, curr, hints := migPipelineSchemas(t)
	ops := Diff(prev, curr, hints)
	withHints := ApplyHintsToOps(ops, hints)
	_ = EmitSQL(withHints, EmitOptions{YongolVersion: "v0.0.0"})
	_ = InferDescription(ops)
	return ops
}
