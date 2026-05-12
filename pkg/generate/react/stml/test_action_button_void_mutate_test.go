//ff:func feature=stml-gen type=test control=sequence
//ff:what NoBodyOps에 포함된 액션이 mutate()를 생성하는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionButtonVoidMutate(t *testing.T) {
	page, _ := stmlparser.ParseReader("delete-page.html", strings.NewReader(`<main>
  <button data-action="DeleteRoom" data-param-room-id="route.RoomID">삭제</button>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		NoBodyOps: map[string]bool{"DeleteRoom": true},
	})
	assertContains(t, code, `deleteRoomMutation.mutate()`)
	assertNotContains(t, code, `mutate({})`)
}
