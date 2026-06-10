//ff:func feature=gen-react type=util control=sequence
//ff:what wrapAuthRetry — withRetry 시 클라이언트 호출식을 withAuthRetry(() => ...) 로 감싼다

package react

// wrapAuthRetry wraps a client call expression in the withAuthRetry
// operation wrapper (emitted by writeRefreshFlow) when withRetry is set.
// The closure re-invokes the whole client call on retry, rebuilding the
// request from args — the openapi-fetch Request itself cannot be re-sent
// after its body is consumed.
func wrapAuthRetry(call string, withRetry bool) string {
	if !withRetry {
		return call
	}
	return "withAuthRetry(() => " + call + ")"
}
