//ff:type feature=crosscheck type=util topic=scenario-check
//ff:what HurlHeader — 요청 헤더 선언

package hurl

// HurlHeader is a request header declaration preceding an HTTP status
// line.
type HurlHeader struct {
	Name  string
	Value string
	Line  int
}
