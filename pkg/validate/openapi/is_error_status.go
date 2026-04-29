//ff:func feature=validate type=util control=sequence dimension=1 topic=response-body-required
//ff:what isErrorStatus — 4xx/5xx 여부 (204/304 예외)

package openapi

import "strconv"

// isErrorStatus returns true when status is a 4xx or 5xx code that requires a
// structured response body under O-5. The 204 No Content and 304 Not Modified
// codes are intentionally bodyless and are excluded. Non-numeric codes (e.g.
// "default", "2XX") fall outside this rule's scope and return false.
func isErrorStatus(status string) bool {
	if status == "204" || status == "304" {
		return false
	}
	n, err := strconv.Atoi(status)
	if err != nil {
		return false
	}
	return n >= 400 && n < 600
}
