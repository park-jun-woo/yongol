//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestByNamePredicates_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	actions := page.Actions
	if !anyActionHasFields(actions) {
		t.Errorf("anyActionHasFields = false")
	}
	if !anyActionHasInputFields(actions) {
		t.Errorf("anyActionHasInputFields = false")
	}
	for _, a := range actions {
		_ = actionHasInputField(a)
	}
	_ = allLoginActions(actions)
	_ = allLoginActions([]stmlparser.ActionBlock{{OperationID: "Login"}})
	cons := byNameConstraints()
	if !anyActionHasZodConstraints(actions, cons) {
		t.Errorf("anyActionHasZodConstraints = false")
	}
	_ = anyActionHasZodConstraints(actions, nil)
}
