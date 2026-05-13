//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=request-id
//ff:what TestRequestIDSource_ContainsKeySymbols — 핵심 심볼 + ulid 의존 + prefix 존재

package middleware

import (
	"strings"
	"testing"
)

func TestRequestIDSource_ContainsKeySymbols(t *testing.T) {
	src := requestIDTypeSource + requestIDMainSource + requestIDFromContextSource + requestIDFromStdContextSource +
		generateRequestIDSource + sanitizeUpstreamIDSource
	for _, must := range []string{
		"func RequestID(",
		"func RequestIDFromContext(",
		"func RequestIDFromStdContext(",
		"github.com/oklog/ulid/v2",
		`const RequestIDPrefix = "req_"`,
		"CtxKeyRequestID",
		"sanitizeUpstreamID",
	} {
		if !strings.Contains(src, must) {
			t.Errorf("requestID sources missing fragment %q", must)
		}
	}
}
