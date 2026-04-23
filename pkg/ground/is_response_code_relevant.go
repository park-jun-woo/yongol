//ff:func feature=rule type=util control=selection
//ff:what isResponseCodeRelevant — status code 첫 글자가 2/4/5 인지 판정
package ground

// isResponseCodeRelevant reports whether the OpenAPI response code should be
// registered by populateResponseSchema. Only 2xx / 4xx / 5xx family codes are
// relevant; 1xx and 3xx are ignored.
func isResponseCodeRelevant(code string) bool {
	switch code[0] {
	case '2', '4', '5':
		return true
	default:
		return false
	}
}
