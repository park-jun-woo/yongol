//ff:type feature=gen-hurl type=model
//ff:what opInfo — operationID에 대한 HTTP method+path 쌍
package hurl

// opInfo holds method and path for an operation.
type opInfo struct {
	Method string
	Path   string
}
