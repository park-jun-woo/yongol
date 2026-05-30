//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestEmitExtPkgImports — 외부 패키지 함수 import 정렬 출력

package ssac

import (
	"strings"
	"testing"
)

func TestEmitExtPkgImports(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		var b strings.Builder
		emitExtPkgImports(&b, map[string]map[string]bool{})
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("SortedPackagesAndFuncs", func(t *testing.T) {
		var b strings.Builder
		emitExtPkgImports(&b, map[string]map[string]bool{
			"zeta":  {"DoThing": true, "AnotherThing": true},
			"alpha": {"Run": true},
		})
		out := b.String()
		want := "from app.services.alpha import run\n" +
			"from app.services.zeta import another_thing, do_thing\n"
		if out != want {
			t.Errorf("got %q, want %q", out, want)
		}
	})
}
