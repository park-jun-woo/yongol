package authz

import (
	"strings"
	"testing"
)

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderAuthzModule_ZeroCov — AuthzModule 소스 생성

func TestRenderAuthzModule_ZeroCov(t *testing.T) {
	out := RenderAuthzModule()
	for _, want := range []string{"export class AuthzModule", "AuthzService", "@Global()"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderAuthzModule missing %q", want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderAuthzService_ZeroCov — AuthzService stub 소스 생성

func TestRenderAuthzService_ZeroCov(t *testing.T) {
	out := RenderAuthzService()
	for _, want := range []string{"export class AuthzService", "async check", "AuthzInput"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderAuthzService missing %q", want)
		}
	}
}
