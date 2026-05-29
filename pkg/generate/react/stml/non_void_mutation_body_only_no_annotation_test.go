//ff:func feature=stml-gen type=test control=sequence
//ff:what body-only(path param 없음) 분기에서 api 함수 직접 참조 형태인지 검증
package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestNonVoidMutationBodyOnlyNoAnnotation(t *testing.T) {
	page, _ := stmlparser.ParseReader("create-room.html", strings.NewReader(`<main>
  <form data-action="CreateRoom">
    <input name="name" />
    <button type="submit">생성</button>
  </form>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
			"CreateRoom": {
				"name": {Type: "string", Required: true},
			},
		},
	})
	assertContains(t, code, `mutationFn: api.CreateRoom`)
	assertNotContains(t, code, `(data) => api.CreateRoom(data)`)
	assertNotContains(t, code, `z.infer<typeof createRoomSchema>)`)
}
