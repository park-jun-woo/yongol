//ff:func feature=gen-gogin type=test control=sequence
//ff:what renderHelperSource 단위 테스트 (annotation + package + body 조립)
package ssac

import (
	"strings"
	"testing"
)

func TestRenderHelperSource(t *testing.T) {
	spec := helperSpec{
		file: "ptr_of.go",
		what: "ptrOf — wraps a value into *T",
		code: "func ptrOf[T any](v T) *T { return &v }\n",
	}
	out := renderHelperSource(spec)

	if !strings.Contains(out, "//ff:what ptrOf — wraps a value into *T") {
		t.Errorf("missing //ff:what annotation:\n%s", out)
	}
	if !strings.Contains(out, "//ff:func") {
		t.Errorf("missing //ff:func annotation:\n%s", out)
	}
	if !strings.Contains(out, "package service\n") {
		t.Errorf("missing package clause:\n%s", out)
	}
	if !strings.HasSuffix(out, spec.code) {
		t.Errorf("body should be appended at the end:\n%s", out)
	}
	// Annotation must precede the package clause.
	if strings.Index(out, "//ff:func") > strings.Index(out, "package service") {
		t.Errorf("annotation should precede package clause")
	}
}
