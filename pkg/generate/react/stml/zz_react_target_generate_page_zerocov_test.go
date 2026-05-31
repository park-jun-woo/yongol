//ff:func feature=stml-gen type=test control=sequence
//ff:what TestReactTargetGeneratePage_ZeroCov — (*ReactTarget).GeneratePage 메서드를 이름으로 직접 호출

package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestReactTargetGeneratePage_ZeroCov(t *testing.T) {
	page, _ := stmlparser.ParseReader("activate-page.html", strings.NewReader(`<main>
  <button data-action="ActivateWorkflow">Activate</button>
</main>`))

	rt := &ReactTarget{}
	code := rt.GeneratePage(page, "", DefaultOptions())
	if code == "" {
		t.Fatal("GeneratePage returned empty code")
	}
	if !strings.Contains(code, "export default function") {
		t.Errorf("generated code missing component export:\n%s", code)
	}
}
