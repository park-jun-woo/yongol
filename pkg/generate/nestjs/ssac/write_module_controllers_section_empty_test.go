//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteModuleControllersSectionEmpty — TestWriteModuleControllersSection — @Module controllers 배열 출력 검증

package ssac

import (
	"strings"
	"testing"
)

func TestWriteModuleControllersSectionEmpty(t *testing.T) {
	var b strings.Builder
	writeModuleControllersSection(&b, nil)
	if b.String() != "  controllers: [\n  ],\n" {
		t.Errorf("got %q", b.String())
	}
}
