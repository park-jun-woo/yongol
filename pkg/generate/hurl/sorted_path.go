//ff:type feature=gen-hurl type=model
//ff:what sortedPath — FK 깊이 정렬용 경로+깊이 쌍
package hurl

// sortedPath holds a path with its FK depth for sorting.
type sortedPath struct {
	Path  string
	Depth int
}
