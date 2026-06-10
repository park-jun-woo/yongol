//ff:type feature=crosscheck type=util topic=scenario-check
//ff:what HurlCapture — [Captures] 라인 (var := jsonpath|header|expr)

package hurl

// HurlCapture is a parsed [Captures] line such as
// `token: jsonpath "$.access_token"` or
// `reqId: header "X-Request-Id"`. Other sources (e.g.
// `csrf: cookie "XSRF-TOKEN"`) are preserved as raw expressions.
type HurlCapture struct {
	Var      string // variable name (e.g. "token")
	Source   string // "jsonpath" | "header" | raw expression prefix
	JSONPath string // non-empty when Source == "jsonpath"
	Header   string // non-empty when Source == "header"
	Line     int    // 1-based line number
}
