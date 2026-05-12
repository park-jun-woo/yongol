//ff:type feature=gen-react type=generator
//ff:what STML 페이지에서 파생된 단일 라우트 정의
package react

// stmlRoute represents a single route derived from an STML page.
type stmlRoute struct {
	Path          string // e.g. "/workflows/:id"
	ComponentName string // e.g. "WorkflowDetail"
	ImportPath    string // e.g. "./pages/workflow-detail"
	Layout        string // layout name (e.g. "app", "auth"); empty = no layout
}
