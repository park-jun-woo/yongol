//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what tm31LinkTargetNotFound — data-link 대상 페이지 부재 ERROR 발화/침묵 검증

package stml_openapi

import "testing"

func TestTM31LinkTargetNotFound(t *testing.T) {
	cases := []TestTM31LinkTargetCase{
		// Target page exists → silent.
		{
			name:      "target_exists",
			html:      `<main><a data-link="building-detail">상세</a></main>`,
			pageNames: []string{"building-detail"},
			wantCount: 0,
		},
		// Typo'd target → ERROR.
		{
			name:      "target_missing",
			html:      `<main><a data-link="bulding-detail">상세</a></main>`,
			pageNames: []string{"building-detail"},
			wantCount: 1,
		},
		// Row link (data-link on the data-each item template) is checked too.
		{
			name: "row_link_target_missing",
			html: `<main><section data-fetch="ListBuildings">
  <ul data-each="buildings">
    <li data-link="nope"><span data-bind="name"></span></li>
  </ul>
</section></main>`,
			pageNames: []string{"building-detail"},
			wantCount: 1,
		},
		// Row-child link inside data-each is checked too.
		{
			name: "row_child_link_target_missing",
			html: `<main><section data-fetch="ListBuildings">
  <ul data-each="buildings">
    <li><a data-link="nope">상세</a></li>
  </ul>
</section></main>`,
			pageNames: []string{"building-detail"},
			wantCount: 1,
		},
		// No links → silent.
		{
			name:      "no_links",
			html:      `<main><h1>hello</h1></main>`,
			pageNames: []string{"building-detail"},
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runTM31LinkTarget(t, c)
		})
	}
}
