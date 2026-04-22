//ff:type feature=gen-react type=model
//ff:what API 엔드포인트 정보를 담는 구조체를 정의한다

package react

// endpoint holds information about a single API endpoint for code generation.
type endpoint struct {
	method     string
	path       string
	opID       string
	pathParams []string // camelCase param names
}
