//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderAuthzModule_ZeroCov — AuthzModule 소스 생성
package authz

import (
	"strings"
	"testing"
)

func TestRenderAuthzModule_ZeroCov(t *testing.T) {
	out := RenderAuthzModule()
	for _, want := range []string{"export class AuthzModule", "AuthzService", "@Global()"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderAuthzModule missing %q", want)
		}
	}
}
