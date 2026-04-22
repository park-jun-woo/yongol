//ff:func feature=rule type=util control=iteration dimension=1
//ff:what hasBearerAuthMiddleware — manifest middleware 목록에 "bearerAuth" 포함 여부 확인
package ground

// hasBearerAuthMiddleware reports whether "bearerAuth" appears in the
// manifest middleware list.
func hasBearerAuthMiddleware(mws []string) bool {
	for _, m := range mws {
		if m == "bearerAuth" {
			return true
		}
	}
	return false
}
