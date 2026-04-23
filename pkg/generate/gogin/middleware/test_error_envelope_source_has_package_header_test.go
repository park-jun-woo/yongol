//ff:func feature=gen-gogin type=test control=sequence topic=error-envelope
//ff:what TestErrorEnvelopeSource_HasPackageHeader — 템플릿에 package 선언 포함

package middleware

import (
	"strings"
	"testing"
)

func TestErrorEnvelopeSource_HasPackageHeader(t *testing.T) {
	if !strings.Contains(errorEnvelopeSource, "package middleware") {
		t.Fatalf("errorEnvelopeSource missing package header")
	}
}
