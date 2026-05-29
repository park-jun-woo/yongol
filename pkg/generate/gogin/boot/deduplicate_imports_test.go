//ff:func feature=gen-gogin type=test control=iteration dimension=2
//ff:what deduplicateImports — 블록별 import 합산 + 중복 제거 + 정렬

package boot

import "testing"

func TestDeduplicateImports(t *testing.T) {
	blocks := []MainBlock{
		{Name: "a", Imports: []string{`"net/http"`, `"github.com/gin-gonic/gin"`}},
		{Name: "b", Imports: []string{`"net/http"`, `"log/slog"`}},
		{Name: "c", Imports: nil},
	}
	got := deduplicateImports(blocks)
	want := []string{`"github.com/gin-gonic/gin"`, `"log/slog"`, `"net/http"`}
	if !equalStrings(got, want) {
		t.Fatalf("deduplicateImports = %v, want %v", got, want)
	}

	if out := deduplicateImports(nil); len(out) != 0 {
		t.Errorf("nil blocks should yield empty, got %v", out)
	}
}
