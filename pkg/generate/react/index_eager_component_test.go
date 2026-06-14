//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what indexEagerComponent — 빈 indexTarget/매칭 경로/불일치 경로별 eager 컴포넌트 판정 검증

package react

import "testing"

func TestIndexEagerComponent(t *testing.T) {
	routes := []stmlRoute{
		{Path: "/", ComponentName: "Home", ImportPath: "./pages/home"},
		{Path: "/workflows/:id", ComponentName: "WorkflowDetail", ImportPath: "./pages/workflow-detail"},
	}

	tests := []struct {
		name        string
		routes      []stmlRoute
		indexTarget string
		want        string
	}{
		{name: "empty indexTarget", routes: routes, indexTarget: "", want: ""},
		{name: "matching root route", routes: routes, indexTarget: "/", want: "Home"},
		{name: "matching nested route", routes: routes, indexTarget: "/workflows/:id", want: "WorkflowDetail"},
		{name: "no matching route", routes: routes, indexTarget: "/missing", want: ""},
		{name: "empty routes", routes: nil, indexTarget: "/", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexEagerComponent(tt.routes, tt.indexTarget); got != tt.want {
				t.Errorf("indexEagerComponent() = %q, want %q", got, tt.want)
			}
		})
	}
}
