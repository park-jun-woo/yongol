//ff:func feature=gen-react-stml type=test control=sequence
//ff:what TestGeneratePageRichZeroCov — fetch/state/each/static/action+form 을 포함한 풍부한 페이지를 GeneratePage 로 렌더해 collect*/render* 다수 커버

package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

const richPageHTML = `<main class="container">
  <section data-fetch="ListReservations" data-param-status="route.Status">
    <ul data-each="reservations">
      <li>
        <span>{item.title}</span>
      </li>
    </ul>
    <button data-action="DeleteReservation" data-param-id="route.ID">삭제</button>
  </section>
  <article data-fetch="GetReservation" data-param-id="route.ID">
    <footer data-state="canCancel">
      <form data-action="UpdateReservation">
        <input data-field="title" />
        <input data-field="note" />
        <button>저장</button>
      </form>
      <button data-action="CancelReservation">취소</button>
    </footer>
  </article>
  <div data-static="info">
    <p>static content</p>
    <button data-action="Login">로그인</button>
  </div>
</main>`

func TestGeneratePageRich_ZeroCov(t *testing.T) {
	page, err := stmlparser.ParseReader("rich.html", strings.NewReader(richPageHTML))
	if err != nil {
		t.Fatalf("ParseReader: %v", err)
	}
	code := GeneratePage(page, "")
	if code == "" {
		t.Fatal("GeneratePage returned empty code")
	}
	// Sanity anchors: mutation hooks and query hooks should be emitted.
	if !strings.Contains(code, "useMutation") {
		t.Errorf("expected a useMutation hook in generated code")
	}
	if !strings.Contains(code, "useQuery") {
		t.Errorf("expected a useQuery hook in generated code")
	}
}

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

func TestTargetMetadata_ZeroCov(t *testing.T) {
	r := &ReactTarget{}
	if r.FileExtension() == "" {
		t.Errorf("FileExtension empty")
	}
	page, _ := stmlparser.ParseReader("rich.html", strings.NewReader(richPageHTML))
	_ = r.Dependencies([]stmlparser.PageSpec{page})
	_ = DefaultOptions()
}
