//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestEmitModelImports — SQLAlchemy 모델 import 정렬 출력

package ssac

import (
	"strings"
	"testing"
)

func TestEmitModelImports(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		var b strings.Builder
		emitModelImports(&b, map[string]bool{})
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("Sorted", func(t *testing.T) {
		var b strings.Builder
		emitModelImports(&b, map[string]bool{"User": true, "Account": true})
		want := "from app.models.models import Account, User\n"
		if b.String() != want {
			t.Errorf("got %q, want %q", b.String(), want)
		}
	})
}
