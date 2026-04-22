//ff:func feature=gen-gogin type=util control=sequence
//ff:what neutralMessage — client-safe default error message per HTTP status code

package ssac

// neutralMessages maps HTTP status code → client-safe default message.
// Used only as a fallback when no explicit message is given in SSaC, so that
// raw internal error text (err.Error()) is never exposed to clients.
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
