//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what tm52DuplicateFieldAcrossForms — 동일 페이지 폼 간 동명 data-field 충돌 WARNING 발화/침묵 검증

package stml_openapi

import "testing"

func TestTM52DuplicateFieldAcrossForms(t *testing.T) {
	cases := []tm52Case{
		{
			// BUG-127 repro: update + create forms nested inside a fetch,
			// both declaring memo. page.Actions is empty here, so only the
			// CollectChildActions(page.Children) walk catches it.
			name: "nested_forms_shared_field",
			html: `<main><article data-fetch="GetBuilding">
  <form data-action="UpdateBuilding">
    <input data-field="memo" />
    <button type="submit">수정</button>
  </form>
  <form data-action="CreateBuilding">
    <input data-field="memo" />
    <button type="submit">생성</button>
  </form>
</article></main>`,
			wantCount: 1,
			wantField: "memo",
		},
		{
			// two top-level forms, distinct field names → no collision.
			name: "distinct_fields",
			html: `<main>
  <form data-action="UpdateBuilding"><input data-field="name" /><button type="submit">a</button></form>
  <form data-action="CreateBuilding"><input data-field="memo" /><button type="submit">b</button></form>
</main>`,
			wantCount: 0,
		},
		{
			// single form → nothing to collide with.
			name:      "single_form",
			html:      `<main><form data-action="UpdateBuilding"><input data-field="memo" /><button type="submit">a</button></form></main>`,
			wantCount: 0,
		},
		{
			// only data-component fields collide → excluded (emit no id).
			name: "component_fields_excluded",
			html: `<main>
  <form data-action="UpdateBuilding"><div data-component="DatePicker" data-field="StartAt" /><button type="submit">a</button></form>
  <form data-action="CreateBuilding"><div data-component="DatePicker" data-field="StartAt" /><button type="submit">b</button></form>
</main>`,
			wantCount: 0,
		},
		{
			// two forms, two shared fields → one warning per field name.
			name: "two_shared_fields",
			html: `<main>
  <form data-action="UpdateBuilding"><input data-field="memo" /><input data-field="is_representative" /><button type="submit">a</button></form>
  <form data-action="CreateBuilding"><input data-field="memo" /><input data-field="is_representative" /><button type="submit">b</button></form>
</main>`,
			wantCount: 2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) { runTM52Case(t, c) })
	}
}
