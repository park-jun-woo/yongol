//ff:func feature=stml-gen type=test control=sequence
//ff:what Delete 액션 버튼에 variant="destructive" 속성 생성을 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestUIButtonDeleteVariant(t *testing.T) {
	page, _ := stmlparser.ParseReader("delete-page.html", strings.NewReader(`<main>
  <button data-action="DeleteRoom" data-param-room-id="route.RoomID">삭제</button>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `variant="destructive"`)
	assertContains(t, code, "<Button ")
}
