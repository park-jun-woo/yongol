//ff:func feature=stml-gen type=test control=sequence
//ff:what 예약 상세 페이지 TSX 생성을 검증
package stml

import ("strings"; "testing"; stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml")

func TestGenerateReservationDetailPage(t *testing.T) {
	page, _ := stmlparser.ParseReader("reservation-detail-page.html", strings.NewReader(`<main class="max-w-2xl mx-auto p-6">
  <article data-fetch="GetReservation" data-param-reservation-id="route.ReservationID">
    <span data-bind="reservation.Status" class="px-3 py-1 text-sm rounded bg-gray-100"></span>
    <dd data-bind="reservation.RoomID" class="font-semibold"></dd>
    <footer data-state="canCancel" class="mt-8 pt-4 border-t">
      <button data-action="CancelReservation" data-param-reservation-id="route.ReservationID">
        예약 취소
      </button>
    </footer>
  </article>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, "export default function ReservationDetailPage()")
	assertContains(t, code, "useParams")
	assertContains(t, code, "ReservationID")
	assertContains(t, code, "api.GetReservation")
	assertContains(t, code, `className="px-3 py-1 text-sm rounded bg-gray-100"`)
	assertContains(t, code, `className="font-semibold"`)
	assertContains(t, code, "<article")
	assertContains(t, code, ".canCancel")
	assertContains(t, code, `className="mt-8 pt-4 border-t"`)
	assertContains(t, code, "<footer")
	assertContains(t, code, "onClick")
	assertContains(t, code, "cancelReservationMutation")
	// isPending on action button
	assertContains(t, code, `disabled={cancelReservationMutation.isPending}`)
	assertContains(t, code, `'처리 중...'`)
}
