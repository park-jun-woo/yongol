//ff:func feature=validate-contract type=util control=selection topic=preserve-safety
//ff:what isResourceAcquireCall — CallExpr 가 close 필요 리소스 반환 호출인지 판정

package contract

import "go/ast"

// isResourceAcquireCall recognises calls whose return value must be
// explicitly Closed. The allowlist covers the handful of stdlib APIs
// that matter for preserved handler bodies:
//
//   os.Open / os.Create / os.OpenFile
//   http.Get / http.Post / client.Do / http.DefaultClient.Do
//   db.Query / db.QueryContext / tx.Query / tx.QueryContext / conn.Query*
//
// db.QueryRow / tx.QueryRow are excluded — Scan releases the row
// automatically, there is no Close obligation on the caller.
func isResourceAcquireCall(call *ast.CallExpr) bool {
	if call == nil {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel == nil {
		return false
	}
	method := sel.Sel.Name
	switch method {
	case "Open", "Create", "OpenFile":
		return leftmostIdentName(sel.X) == "os"
	case "Get", "Post", "PostForm", "Head":
		return leftmostIdentName(sel.X) == "http"
	case "Do":
		return true
	case "Query", "QueryContext":
		return true
	}
	return false
}
