//ff:func feature=stml-gen type=test control=sequence
//ff:what 예약 목록 페이지 TSX 생성을 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGenerateMyReservationsPage(t *testing.T) {
	page, _ := stmlparser.ParseReader("my-reservations-page.html", strings.NewReader(`<main class="max-w-4xl mx-auto p-6">
  <section data-fetch="ListMyReservations" class="mb-8">
    <ul data-each="reservations" class="space-y-3">
      <li class="flex justify-between p-4 border rounded">
        <span data-bind="RoomID" class="font-semibold"></span>
        <span data-bind="Status" class="px-2 py-1 text-sm rounded bg-gray-100"></span>
      </li>
    </ul>
    <p data-state="reservations.empty" class="text-gray-400">예약이 없습니다</p>
  </section>
  <div data-action="CreateReservation" class="space-y-4">
    <input data-field="RoomID" type="number" placeholder="스터디룸 번호" class="w-full px-3 py-2 border rounded" />
    <div data-component="DatePicker" data-field="StartAt" />
    <div data-component="DatePicker" data-field="EndAt" />
    <button type="submit">예약하기</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, "export default function MyReservationsPage()")
	assertContains(t, code, `<main className="mx-auto max-w-4xl px-4 py-8 space-y-6">`)
	assertContains(t, code, `<h1 className="text-2xl font-bold">My Reservations</h1>`)
	assertContains(t, code, "useQuery")
	assertContains(t, code, "api.ListMyReservations")
	assertContains(t, code, `className="mb-8"`)
	assertContains(t, code, `className="space-y-3"`)
	assertContains(t, code, "<Table")
	assertContains(t, code, "<THead>")
	assertContains(t, code, "<TBody>")
	assertContains(t, code, "<TR")
	assertContains(t, code, "<TH>Room ID</TH>")
	assertContains(t, code, "<TH>Status</TH>")
	assertContains(t, code, "<TD>{item.RoomID}</TD>")
	assertContains(t, code, "<TD>{item.Status}</TD>")
	assertContains(t, code, "예약이 없습니다")
	assertContains(t, code, "length === 0")
	assertContains(t, code, "api.CreateReservation")
	assertContains(t, code, `placeholder="스터디룸 번호"`)
	assertContains(t, code, "'예약하기'")
	assertContains(t, code, "import { Table, THead, TBody, TR, TH, TD } from '@/components/ui/Table'")
	assertContains(t, code, "import DatePicker from '@/components/DatePicker'")
	assertContains(t, code, "<DatePicker")
	assertContains(t, code, "queryKey: ['ListMyReservations']")
	// label + id attributes on form fields
	assertContains(t, code, `<label htmlFor="RoomID">Room ID</label>`)
	assertContains(t, code, `id="RoomID"`)
	// isPending on submit button
	assertContains(t, code, `disabled={createReservationMutation.isPending}`)
	assertContains(t, code, `'처리 중...'`)
}
