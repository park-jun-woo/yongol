//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameEachKeyFields_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	raif := map[string]map[string]map[string]bool{
		"ListItems": {"items": {"id": true, "name": true}},
		"ListSub":   {"x": {"id": true}},
	}
	populateEachKeyFields(&page, raif)
	populateEachKeyFields(&page, nil)
	populateEachKeyFieldsInChildren(page.Children, raif)
	for i := range page.Children {
		populateEachKeyFieldForChild(&page.Children[i], raif)
	}

	f := page.Fetches[0]
	setEachKeyFieldsInFetch(&f, "ListItems", raif)
	setEachKeyFieldsInChildren(f.Children, "ListItems", raif)
	for i := range f.Children {
		setEachKeyFieldForChild(&f.Children[i], "ListItems", raif)
	}
	if len(f.Eaches) > 0 {
		setKeyFieldIfHasID(&f.Eaches[0], "ListItems", raif)
		setKeyFieldIfHasID(&f.Eaches[0], "Missing", raif)
	}
}
