//ff:type feature=crosscheck type=util topic=scenario-check
//ff:what Hurl 요청/응답 쌍 타입 정의
package hurl

// HurlEntry represents one request/response pair extracted from a .hurl file.
type HurlEntry struct {
	Method     string
	Path       string
	StatusCode string
	File       string
	Line       int
}
