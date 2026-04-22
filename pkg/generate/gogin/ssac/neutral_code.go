//ff:func feature=gen-gogin type=util control=sequence
//ff:what neutralCode — HTTP status code 별 machine-readable 에러 코드 기본값

package ssac

// neutralCodes maps HTTP status code → default machine-readable error code.
// Paired with neutralMessages: both are looked up when SSaC does not supply
// an explicit code / message attribute, so the emitted handler produces a
// consistent pair like {error: "Payment required", code: "payment_required"}.
var neutralCodes = map[int]string{
	400: "bad_request",
	401: "unauthorized",
	402: "payment_required",
	403: "forbidden",
	404: "not_found",
	409: "conflict",
	422: "unprocessable_entity",
	429: "too_many_requests",
	500: "internal_error",
	502: "bad_gateway",
	503: "service_unavailable",
}

// neutralCode returns the default error code for status.
// Falls back to "internal_error" when status is unknown.
func neutralCode(status int) string {
	if c, ok := neutralCodes[status]; ok {
		return c
	}
	return "internal_error"
}
