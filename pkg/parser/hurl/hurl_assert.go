//ff:type feature=crosscheck type=util topic=scenario-check
//ff:what HurlAssert — `jsonpath "$.path" ...` 어서션 라인

package hurl

// HurlAssert is a parsed `jsonpath "$.path" ...` assertion line.
type HurlAssert struct {
	JSONPath string // e.g. "$.user.id"
	Line     int    // 1-based line number in the source file
}
