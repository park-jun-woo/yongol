//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderAuthzModule_ZeroCov — AuthzModule 소스 생성
package authz

import (
	"strings"
	"testing"
)

func TestRenderAuthzService_ZeroCov(t *testing.T) {
	out := RenderAuthzService()
	for _, want := range []string{"export class AuthzService", "async check", "AuthzInput"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderAuthzService missing %q", want)
		}
	}
}
