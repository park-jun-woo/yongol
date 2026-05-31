//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestNestDependencies_ZeroCov — 런타임 의존성 맵
package scaffold

import (
	"strings"
	"testing"
)

func TestRenderPackageJSON_ZeroCov(t *testing.T) {
	if _, err := RenderPackageJSON(""); err == nil {
		t.Error("expected error for empty projectID")
	}
	out, err := RenderPackageJSON("myapp")
	if err != nil {
		t.Fatalf("RenderPackageJSON error: %v", err)
	}
	for _, want := range []string{`"name": "myapp"`, "nest build", "@nestjs/core", "prisma"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderPackageJSON missing %q", want)
		}
	}
}
