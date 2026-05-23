//ff:type feature=cli-init type=model
//ff:what httpRoute — features.yaml path 필드에서 파싱된 HTTP 메서드+URI

package cliinit

// httpRoute holds the parsed method and URI from a features.yaml path field.
type httpRoute struct {
	Method string // lowercase: "get", "post", "put", "patch", "delete"
	URI    string // e.g. "/workflows/{id}"
}
