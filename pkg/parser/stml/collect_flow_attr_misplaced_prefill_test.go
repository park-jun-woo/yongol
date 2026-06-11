//ff:func feature=stml-parse type=test control=sequence
//ff:what TestCollectFlowAttrMisplacedPrefill — data-action 밖의 data-prefill이 misplaced로 기록되는지 검증

package stml

import (
	"strings"
	"testing"
)

func TestCollectFlowAttrMisplacedPrefill(t *testing.T) {
	// data-prefill on a non-action element is misplaced (same rule as
	// data-capture/data-redirect — valid only on the data-action element).
	input := `<main>
  <div data-prefill="GetRule"><span>x</span></div>
  <form data-action="UpdateRule" data-prefill="GetRule"><button type="submit">go</button></form>
</main>`

	page, _ := ParseReader("p.html", strings.NewReader(input))

	if len(page.FlowAttrMisplaced) != 1 {
		t.Fatalf("FlowAttrMisplaced = %+v, want exactly one record", page.FlowAttrMisplaced)
	}
	got := page.FlowAttrMisplaced[0]
	if got.Attr != "data-prefill" || got.Tag != "div" {
		t.Errorf("misplaced = %+v, want {data-prefill div}", got)
	}
}
