//ff:func feature=stml-parse type=parser control=sequence
//ff:what TestPhase4_StateText — state element tag, className, and inner text extraction

package stml

import (
	"strings"
	"testing"
)

func TestPhase4_StateText(t *testing.T) {
	input := `<main>
  <section data-fetch="ListMyReservations">
    <p data-state="reservations.empty" class="text-gray-400">예약이 없습니다</p>
  </section>
</main>`

	page, diags := ParseReader("test.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	state := page.Fetches[0].States[0]
	if state.Tag != "p" {
		t.Errorf("State.Tag = %q, want %q", state.Tag, "p")
	}
	if state.ClassName != "text-gray-400" {
		t.Errorf("State.ClassName = %q, want %q", state.ClassName, "text-gray-400")
	}
	if state.Text != "예약이 없습니다" {
		t.Errorf("State.Text = %q, want %q", state.Text, "예약이 없습니다")
	}
}
