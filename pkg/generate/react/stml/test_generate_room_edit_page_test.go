//ff:func feature=stml-gen type=test control=sequence
//ff:what 스터디룸 수정 페이지 TSX 생성을 검증
package stml

import ("strings"; "testing"; stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml")

func TestGenerateRoomEditPage(t *testing.T) {
	page, _ := stmlparser.ParseReader("room-edit-page.html", strings.NewReader(`<main class="max-w-2xl mx-auto p-6">
  <div data-action="UpdateRoom" data-param-room-id="route.RoomID" class="space-y-4">
    <input data-field="Name" placeholder="스터디룸 이름" class="w-full px-3 py-2 border rounded" />
    <input data-field="Capacity" type="number" placeholder="수용 인원" class="w-full px-3 py-2 border rounded" />
    <input data-field="Location" placeholder="위치" class="w-full px-3 py-2 border rounded" />
    <button type="submit">수정</button>
  </div>
  <footer data-state="canDelete" class="mt-8 pt-4 border-t">
    <button data-action="DeleteRoom" data-param-room-id="route.RoomID">
      스터디룸 삭제
    </button>
  </footer>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, "export default function RoomEditPage()")
	assertContains(t, code, "useParams")
	assertContains(t, code, "RoomID")
	assertContains(t, code, "api.UpdateRoom")
	assertContains(t, code, "api.DeleteRoom")
	assertContains(t, code, `placeholder="스터디룸 이름"`)
	assertContains(t, code, `placeholder="수용 인원"`)
	assertContains(t, code, `placeholder="위치"`)
	assertContains(t, code, "'수정'")
	assertContains(t, code, "onClick")
	assertContains(t, code, "deleteRoomMutation.mutate")
	assertNotContains(t, code, "deleteRoomForm")
	// label + id on form fields
	assertContains(t, code, `<label htmlFor="Name">Name</label>`)
	assertContains(t, code, `<label htmlFor="Capacity">Capacity</label>`)
	assertContains(t, code, `<label htmlFor="Location">Location</label>`)
	// isPending on submit and action buttons
	assertContains(t, code, `disabled={updateRoomMutation.isPending}`)
	assertContains(t, code, `disabled={deleteRoomMutation.isPending}`)
}
