//ff:func feature=gen-gogin type=test control=sequence topic=request-id
//ff:what request_id_source 스냅샷 — 핵심 심볼 + ulid 의존 + prefix 존재

package middleware

import (
	"strings"
	"testing"
)

func TestRequestIDSource_ContainsKeySymbols(t *testing.T) {
	src := requestIDSource
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
			t.Errorf("requestIDSource missing fragment %q", must)
		}
	}
}

func TestRequestIDSource_HasPackageHeader(t *testing.T) {
	if !strings.Contains(requestIDSource, "package middleware") {
		t.Fatalf("requestIDSource missing package header")
	}
}
