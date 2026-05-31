//ff:func feature=stml-gen type=test control=sequence
//ff:what TestGeneratePageRichZeroCov — fetch/state/each/static/action+form 을 포함한 풍부한 페이지를 GeneratePage 로 렌더해 collect*/render* 다수 커버
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePageRich_ZeroCov(t *testing.T) {
	page, err := stmlparser.ParseReader("rich.html", strings.NewReader(richPageHTML))
	if err != nil {
		t.Fatalf("ParseReader: %v", err)
	}
	code := GeneratePage(page, "")
	if code == "" {
		t.Fatal("GeneratePage returned empty code")
	}
	// Sanity anchors: mutation hooks and query hooks should be emitted.
	if !strings.Contains(code, "useMutation") {
		t.Errorf("expected a useMutation hook in generated code")
	}
	if !strings.Contains(code, "useQuery") {
		t.Errorf("expected a useQuery hook in generated code")
	}
}
