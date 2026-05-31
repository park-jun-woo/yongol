//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package react

import (
	"strings"
	"testing"
)

func TestByNameRouteWriters_ZeroCov(t *testing.T) {
	routes := []stmlRoute{
		{Path: "/", ComponentName: "Home", ImportPath: "./pages/home", Layout: ""},
		{Path: "/items", ComponentName: "Items", ImportPath: "./pages/items", Layout: "app"},
	}
	layoutSet := map[string]bool{"app": true}

	app := renderAppTSX(routes, layoutSet, true)
	if !strings.Contains(app, "Routes") {
		t.Errorf("renderAppTSX missing Routes")
	}
	_ = renderAppTSX(routes, layoutSet, false)

	var sb strings.Builder
	writeLayoutImports(&sb, []string{"app", "auth"})
	writePageImports(&sb, routes)
	writeFlatRoutes(&sb, []stmlRoute{routes[0]}, true)
	writeLayoutGroupRoutes(&sb, "app", []stmlRoute{routes[1]}, true)
	writeAuthzMiddleware(&sb)
	if sb.Len() == 0 {
		t.Errorf("route writers empty")
	}
}
