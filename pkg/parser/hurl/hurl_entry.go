//ff:type feature=crosscheck type=util topic=scenario-check
//ff:what HurlEntry — Hurl 요청/응답 쌍 (XOH-01~09 body/headers/captures/asserts 포함)

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

	// URLVar holds the name of the {{var}} placeholder used as the URL
	// prefix on the request line (e.g. "host", "authurl", "rest"). It is
	// "" when the request line uses an absolute http(s):// URL. Used by
	// XOH-01 to skip OpenAPI path matching for external services
	// (URLVar != "" && URLVar != "host").
	URLVar string

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
