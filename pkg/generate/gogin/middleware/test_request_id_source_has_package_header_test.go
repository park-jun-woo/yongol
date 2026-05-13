//ff:func feature=gen-gogin type=test control=sequence topic=request-id
//ff:what TestRequestIDSource_HasPackageHeader — 템플릿에 package middleware 포함

package middleware

import (
	"strings"
	"testing"
)

func TestRequestIDSource_HasPackageHeader(t *testing.T) {
	if !strings.Contains(requestIDMainSource, "package middleware") {
		t.Fatalf("requestIDMainSource missing package header")
	}
}
