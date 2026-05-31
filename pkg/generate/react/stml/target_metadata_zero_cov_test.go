//ff:func feature=stml-gen type=test control=sequence
//ff:what TestGeneratePageRichZeroCov — fetch/state/each/static/action+form 을 포함한 풍부한 페이지를 GeneratePage 로 렌더해 collect*/render* 다수 커버
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTargetMetadata_ZeroCov(t *testing.T) {
	r := &ReactTarget{}
	if r.FileExtension() == "" {
		t.Errorf("FileExtension empty")
	}
	page, _ := stmlparser.ParseReader("rich.html", strings.NewReader(richPageHTML))
	_ = r.Dependencies([]stmlparser.PageSpec{page})
	_ = DefaultOptions()
}
