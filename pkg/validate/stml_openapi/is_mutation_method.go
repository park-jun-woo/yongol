//ff:func feature=validate type=util control=selection topic=stml-openapi
//ff:what isMutationMethod — HTTP 메서드가 상태 변경 mutation(POST/PUT/PATCH/DELETE)인지 판정

package stml_openapi

// isMutationMethod reports whether an HTTP method is a state-changing mutation
// (POST/PUT/PATCH/DELETE). GET (and any other method) is non-mutating.
// Extracted from tm57MutationRedirectRequired so that rule stays a flat guard
// chain.
func isMutationMethod(method string) bool {
	switch method {
	case "POST", "PUT", "PATCH", "DELETE":
		return true
	default:
		return false
	}
}
