//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"testing"
)

func TestFormatColumnList(t *testing.T) {
	if got := formatColumnList([]string{"a", "b"}); got != `"a", "b"` {
		t.Errorf("unexpected: %q", got)
	}
	if got := formatColumnList([]string{"only"}); got != `"only"` {
		t.Errorf("unexpected single: %q", got)
	}
	if got := formatColumnList(nil); got != "" {
		t.Errorf("expected empty for nil, got %q", got)
	}
}
