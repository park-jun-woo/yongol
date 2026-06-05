//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteModuleImportsSectionFull — TestWriteModuleImportsSection — @Module imports 배열(조건부 Queue/Authz/cross-feature) 출력 검증

package ssac

import (
	"strings"
	"testing"
)

func TestWriteModuleImportsSectionFull(t *testing.T) {
	var b strings.Builder
	deps := moduleDeps{
		NeedsQueue:    true,
		NeedsAuthz:    true,
		CrossFeatures: []string{"billing"},
	}
	writeModuleImportsSection(&b, deps)

	want := "  imports: [\n" +
		"    PrismaModule,\n" +
		"    QueueModule,\n" +
		"    AuthzModule,\n" +
		"    BillingModule,\n" +
		"  ],\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
