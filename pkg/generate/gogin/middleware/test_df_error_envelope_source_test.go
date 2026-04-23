//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=error-envelope
//ff:what TestErrorEnvelopeSource_ContainsKeySymbols — envelope 구조체/함수 심볼 존재 확인

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
