//ff:func feature=stml-gen type=test control=sequence
//ff:what pageCrumbField — 등재+fetch 보유 시 필드 / fetch 부재·미등재·nil 맵 빈 문자열 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestPageCrumbField(t *testing.T) {
	fields := map[string]string{"building-detail": "building_name"}
	withFetch := stmlparser.PageSpec{Name: "building-detail", Fetches: []stmlparser.FetchBlock{{OperationID: "GetBuilding"}}}

	t.Run("listed page with a fetch yields the field", func(t *testing.T) {
		if got := pageCrumbField(withFetch, fields); got != "building_name" {
			t.Errorf("pageCrumbField = %q, want building_name", got)
		}
	})

	t.Run("fetch-less page yields empty (TM-50 territory)", func(t *testing.T) {
		if got := pageCrumbField(stmlparser.PageSpec{Name: "building-detail"}, fields); got != "" {
			t.Errorf("pageCrumbField = %q, want empty", got)
		}
	})

	t.Run("unlisted page and nil map yield empty", func(t *testing.T) {
		other := stmlparser.PageSpec{Name: "home", Fetches: []stmlparser.FetchBlock{{OperationID: "ListItems"}}}
		if got := pageCrumbField(other, fields); got != "" {
			t.Errorf("pageCrumbField = %q, want empty for an unlisted page", got)
		}
		if got := pageCrumbField(withFetch, nil); got != "" {
			t.Errorf("pageCrumbField = %q, want empty for a nil map", got)
		}
	})
}
