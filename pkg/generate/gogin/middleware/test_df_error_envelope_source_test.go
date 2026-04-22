//ff:func feature=gen-gogin type=test control=sequence topic=error-envelope
//ff:what error_envelope_source 스냅샷 — envelope 구조체 + status 매핑 + 미들웨어 심볼

package middleware

import (
	"strings"
	"testing"
)

func TestErrorEnvelopeSource_ContainsKeySymbols(t *testing.T) {
	src := errorEnvelopeSource
	for _, must := range []string{
		"type ErrorEnvelope struct",
		`Error       string`,
		`Message     string`,
		`RequestID   string`,
		`RetryAfter  int64`,
		`Limit       int64`,
		`FieldErrors map[string]string`,
		"func DefaultCodeFor(",
		"func DefaultMessageFor(",
		"func WriteEnvelope(",
		"func WriteEnvelopeWithContext(",
		"func ErrorEnvelopeMiddleware()",
		`"rate_limit_exceeded"`,
		`"payload_too_large"`,
		`"validation_failed"`,
		`"internal_error"`,
	} {
		if !strings.Contains(src, must) {
			t.Errorf("errorEnvelopeSource missing fragment %q", must)
		}
	}
}

func TestErrorEnvelopeSource_HasPackageHeader(t *testing.T) {
	if !strings.Contains(errorEnvelopeSource, "package middleware") {
		t.Fatalf("errorEnvelopeSource missing package header")
	}
}
