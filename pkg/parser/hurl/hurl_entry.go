//ff:type feature=crosscheck type=util topic=scenario-check
//ff:what Hurl 요청/응답 쌍 타입 정의 (XOH-01~09 용 captures/asserts/body/headers 포함)
package hurl

// HurlEntry represents one request/response pair extracted from a .hurl file.
//
// Fields beyond Method/Path/StatusCode (BodyFields / Asserts / Captures /
// Headers) are populated by Phase002's extended parser and consumed by
// the XOH-01~09 cross-check rules. Rules that only need method/path/status
// continue to work unchanged.
type HurlEntry struct {
	Method     string
	Path       string
	StatusCode string
	File       string
	Line       int

	// BodyFields lists top-level JSON request body field names found
	// between the request line and the HTTP status line. Used by XOH-03
	// (request body field in OpenAPI schema).
	BodyFields []string

	// Asserts lists jsonpath assertions under the [Asserts] section
	// (also captured when `jsonpath "$..."` appears bare on an assert
	// line). Used by XOH-04 (assert jsonpath reachable in response).
	Asserts []HurlAssert

	// Captures lists [Captures] entries (var := source expression).
	// Used by XOH-08 (capture reachable in response) and XOH-09
	// (unused capture).
	Captures []HurlCapture

	// Headers lists request header names in their declared order. Used
	// by XOH-07 (CSRF header on cookie-mode mutation) and XOH-06 (auth
	// precondition via Cookie / Authorization header).
	Headers []HurlHeader
}

// HurlAssert is a parsed `jsonpath "$.path" ...` assertion line.
type HurlAssert struct {
	JSONPath string // e.g. "$.user.id"
	Line     int    // 1-based line number in the source file
}

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

// HurlHeader is a request header declaration preceding an HTTP status
// line.
type HurlHeader struct {
	Name  string
	Value string
	Line  int
}
