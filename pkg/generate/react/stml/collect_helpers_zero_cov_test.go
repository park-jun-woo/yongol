//ff:func feature=stml-gen type=test control=sequence
//ff:what TestGeneratePageRichZeroCov — fetch/state/each/static/action+form 을 포함한 풍부한 페이지를 GeneratePage 로 렌더해 collect*/render* 다수 커버
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectHelpers_ZeroCov(t *testing.T) {
	page, err := stmlparser.ParseReader("rich.html", strings.NewReader(richPageHTML))
	if err != nil {
		t.Fatalf("ParseReader: %v", err)
	}
	if len(collectAllActions(page.Children)) == 0 {
		t.Errorf("collectAllActions returned none")
	}
	_ = collectAllParams(page)
}
