//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameGeneratePage_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	out := GeneratePage(page, "", GenerateOptions{
		RequestConstraints:      byNameConstraints(),
		ResponseArrayItemFields: map[string]map[string]map[string]bool{"ListItems": {"items": {"id": true}}},
	})
	if out == "" {
		t.Errorf("GeneratePage empty")
	}
}
