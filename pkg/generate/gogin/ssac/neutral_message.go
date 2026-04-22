//ff:func feature=gen-gogin type=util control=sequence
//ff:what neutralMessage — HTTP status code 별 중립 에러 메시지

package ssac

// neutralMessages maps HTTP status code → client-safe default message.
// SSaC 에 명시 메시지가 없는 경우에만 fallback 으로 사용. 내부 에러 원문
// (err.Error()) 을 외부에 노출하지 않기 위한 기반.
var neutralMessages = map[int]string{
	400: "Bad request",
	401: "Unauthorized",
	402: "Payment required",
	403: "Forbidden",
	404: "Not found",
	409: "Conflict",
	422: "Unprocessable entity",
	429: "Too many requests",
	500: "Internal error",
	502: "Bad gateway",
	503: "Service unavailable",
}

// neutralMessage returns the default neutral message for status.
// Falls back to "Internal error" when status is unknown.
func neutralMessage(status int) string {
	if m, ok := neutralMessages[status]; ok {
		return m
	}
	return "Internal error"
}
