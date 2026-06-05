//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteModuleImportsSectionMinimal — TestWriteModuleImportsSection — @Module imports 배열(조건부 Queue/Authz/cross-feature) 출력 검증

package ssac

import (
	"strings"
	"testing"
)

func TestWriteModuleImportsSectionMinimal(t *testing.T) {
	var b strings.Builder
	writeModuleImportsSection(&b, moduleDeps{})

	want := "  imports: [\n    PrismaModule,\n  ],\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
