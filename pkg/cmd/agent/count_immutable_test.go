//ff:func feature=agent type=test control=sequence
//ff:what TestCountImmutable — countImmutable 이 immutable 파일을 가리키는 ERROR 만 세는지 검증
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCountImmutable(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Level: diagnostic.LevelError, File: "specs/features.yaml"},
		{Level: diagnostic.LevelError, File: "specs/tests/x.hurl"},
		{Level: diagnostic.LevelError, File: "specs/openapi.yaml"}, // mutable
		{Level: diagnostic.LevelWarning, File: "features.yaml"},    // not ERROR
	}
	if got := countImmutable(diags); got != 2 {
		t.Errorf("countImmutable = %d, want 2", got)
	}
	if got := countImmutable(nil); got != 0 {
		t.Errorf("countImmutable(nil) = %d, want 0", got)
	}
}
