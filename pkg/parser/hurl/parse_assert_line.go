//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what parseAssertLine — `jsonpath "$.path" isXxx` 라인에서 jsonpath 추출

package hurl

// parseAssertLine extracts the first jsonpath argument in a line like
// `jsonpath "$.user.id" isInteger`. Returns false when the line lacks a
// jsonpath expression.
func parseAssertLine(line string, lineNum int) (HurlAssert, bool) {
	m := reJSONPath.FindStringSubmatch(line)
	if m == nil {
		return HurlAssert{}, false
	}
	return HurlAssert{JSONPath: m[1], Line: lineNum}, true
}
