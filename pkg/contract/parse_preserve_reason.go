//ff:func feature=contract type=util control=sequence
//ff:what ParsePreserveReason — //ff:preserve reason="..." 주석에서 reason 문자열을 추출

package contract

import (
	"os"
	"regexp"
)

// preserveReasonRe captures the quoted reason value on a
// `//ff:preserve reason="..."` annotation. Double quotes inside the
// reason are not expected; callers that need escaping should revisit
// this regex along with the annotation spec.
var preserveReasonRe = regexp.MustCompile(`(?m)^\s*//ff:preserve\b[^\n]*\breason="([^"]*)"`)

// ParsePreserveReason reads filePath and returns the `reason` value
// from a `//ff:preserve reason="..."` annotation, or the empty string
// when no such annotation exists. Read errors are returned as-is so
// callers can distinguish missing files from missing annotations.
func ParsePreserveReason(filePath string) (string, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	m := preserveReasonRe.FindSubmatch(src)
	if m == nil {
		return "", nil
	}
	return string(m[1]), nil
}
