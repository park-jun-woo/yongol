//ff:type feature=crosscheck type=util topic=scenario-check
//ff:what HurlCapture — [Captures] 라인 (var := jsonpath|header|expr)

package hurl

// HurlCapture is a parsed [Captures] line such as
// `token: jsonpath "$.access_token"` or
// `csrf: header "X-CSRF-Token"`.
type HurlCapture struct {
	Var      string // variable name (e.g. "token")
	Source   string // "jsonpath" | "header" | raw expression prefix
	JSONPath string // non-empty when Source == "jsonpath"
	Header   string // non-empty when Source == "header"
	Line     int    // 1-based line number
}
