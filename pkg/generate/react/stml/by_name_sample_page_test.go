//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func byNameSamplePage(t *testing.T) stmlparser.PageSpec {
	t.Helper()
	const src = `<main data-layout="app" data-route="/items">
  <section data-fetch="ListItems" data-param-status="route.status">
    <ul data-each="items">
      <li data-bind="name"></li>
      <span data-component="Badge" data-bind="status"></span>
    </ul>
    <span data-bind="total"></span>
    <p data-state="items.empty">없음</p>
    <section data-fetch="ListSub"><span data-bind="x"></span></section>
    <h2>Header</h2>
  </section>
  <div data-action="CreateItem" data-param-id="route.id">
    <input data-field="Name" type="text" />
    <input data-field="Count" type="number" />
    <span data-component="DatePicker" data-field="Due"></span>
    <button type="submit">Create</button>
  </div>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <input data-field="Password" type="password" />
    <button type="submit">Login</button>
  </div>
</main>`
	page, diags := stmlparser.ParseReader("items-page.html", strings.NewReader(src))
	if len(diags) > 0 {
		t.Fatalf("parse diags: %v", diags)
	}
	return page
}
