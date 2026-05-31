//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderPyproject — pyproject.toml / requirements.txt / 의존성 목록 생성 검증
package scaffold

import (
	"testing"
)

func TestRenderPyprojectEmptyID(t *testing.T) {
	if _, err := RenderPyproject(""); err == nil {
		t.Error("expected error for empty projectID")
	}
}
