//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — react 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestByNamePureHelpers_ZeroCov(t *testing.T) {
	if got := buildCNArgs("base", []string{"v"}, []string{"s"}); len(got) == 0 {
		t.Errorf("buildCNArgs empty")
	}
	_ = buildCNArgs("", nil, nil)
	if got := buildDestructParams([]string{"v"}, []string{"s"}, "default", "md"); len(got) == 0 {
		t.Errorf("buildDestructParams empty")
	}
	_ = buildDestructParams(nil, nil, "", "")

	if got := extractPathParams("/api/items/{id}/sub/{subId}"); len(got) != 2 {
		t.Errorf("extractPathParams = %v", got)
	}
	_ = extractPathParams("/api/items")

	for _, tag := range []string{"button", "input", "select", "textarea", "form", "table", "label", "a", "div"} {
		_ = htmlAttrsType(tag)
		_ = inferHTMLElement(tag)
	}
	for _, name := range []string{"Button", "Input", "Select", "Textarea", "Form", "Table", "Label", "Link", "Card"} {
		_ = inferHTMLTag(name)
	}
	dspec := &design.DesignSpec{Colors: map[string]string{"primary": "#fff"}}
	if got := designColor(dspec, "primary", "def"); got != "#fff" {
		t.Errorf("designColor = %q", got)
	}
	_ = designColor(dspec, "missing", "def")
	_ = designColor(nil, "primary", "def")

	set := buildLayoutSet([]stml.LayoutSpec{{Name: "app"}, {Name: "auth"}})
	if !set["app"] {
		t.Errorf("buildLayoutSet missing app")
	}

	grp := groupRoutesByLayout([]stmlRoute{{Path: "/a", Layout: "app"}, {Path: "/b", Layout: "app"}})
	if len(grp["app"]) != 2 {
		t.Errorf("groupRoutesByLayout = %v", grp)
	}
}
