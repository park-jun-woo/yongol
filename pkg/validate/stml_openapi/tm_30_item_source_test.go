//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what tm30ItemSource — item.* 소스의 each 외부 사용·item 스키마 부재 ERROR 발화/침묵 검증

package stml_openapi

import "testing"

func TestTM30ItemSource(t *testing.T) {
	raif := map[string]map[string]map[string]bool{
		"GetUnit": {"photos": {"id": true, "photo": true}},
	}
	cases := []TestTM30ItemSourceCase{
		// item field exists in the each item schema → silent.
		{
			name: "valid_item_field",
			html: `<main><section data-fetch="GetUnit">
  <ul data-each="photos"><li>
    <span data-bind="caption"></span>
    <button data-action="DeletePhoto" data-param-photo-id="item.id">x</button>
  </li></ul>
</section></main>`,
			raif:      raif,
			wantCount: 0,
		},
		// item field missing from the item schema → ERROR.
		{
			name: "missing_item_field",
			html: `<main><section data-fetch="GetUnit">
  <ul data-each="photos"><li>
    <span data-bind="caption"></span>
    <button data-action="DeletePhoto" data-param-photo-id="item.nope">x</button>
  </li></ul>
</section></main>`,
			raif:      raif,
			wantCount: 1,
		},
		// dotted source resolves its first segment against the schema.
		{
			name: "dotted_first_segment",
			html: `<main><section data-fetch="GetUnit">
  <ul data-each="photos"><li>
    <span data-bind="caption"></span>
    <button data-action="DeletePhoto" data-param-photo-id="item.photo.id">x</button>
  </li></ul>
</section></main>`,
			raif:      raif,
			wantCount: 0,
		},
		// item.* on a page-level action (outside any data-each) → ERROR.
		{
			name:      "outside_each_page_action",
			html:      `<main><button data-action="DeletePhoto" data-param-photo-id="item.id">x</button></main>`,
			raif:      raif,
			wantCount: 1,
		},
		// item.* on a fetch-internal action outside data-each → ERROR.
		{
			name: "outside_each_fetch_action",
			html: `<main><section data-fetch="GetUnit">
  <span data-bind="caption"></span>
  <button data-action="DeletePhoto" data-param-photo-id="item.id">x</button>
</section></main>`,
			raif:      raif,
			wantCount: 1,
		},
		// item.* on a data-fetch param → ERROR (a fetch is never row-scoped).
		{
			name:      "fetch_param_item_source",
			html:      `<main><section data-fetch="GetUnit" data-param-photo-id="item.id"><span data-bind="caption"></span></section></main>`,
			raif:      raif,
			wantCount: 1,
		},
		// unresolved item schema (op/each field unknown) → silent (TM-01/07).
		{
			name: "unresolved_schema_silent",
			html: `<main><section data-fetch="GetUnit">
  <ul data-each="photos"><li>
    <span data-bind="caption"></span>
    <button data-action="DeletePhoto" data-param-photo-id="item.id">x</button>
  </li></ul>
</section></main>`,
			raif:      map[string]map[string]map[string]bool{},
			wantCount: 0,
		},
		// route.* sources never fire TM-30.
		{
			name: "route_sources_silent",
			html: `<main><section data-fetch="GetUnit" data-param-building-id="route.BuildingID">
  <ul data-each="photos"><li>
    <span data-bind="caption"></span>
    <button data-action="DeletePhoto" data-param-photo-id="route.PhotoID">x</button>
  </li></ul>
</section></main>`,
			raif:      raif,
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runTM30ItemSource(t, c)
		})
	}
}
