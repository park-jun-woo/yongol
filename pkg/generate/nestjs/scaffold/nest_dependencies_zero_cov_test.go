//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestDependencies_ZeroCov — 런타임 의존성 맵
package scaffold

import (
	"testing"
)

func TestNestDependencies_ZeroCov(t *testing.T) {
	deps := nestDependencies()
	if deps["@nestjs/core"] == "" {
		t.Error("expected @nestjs/core dependency")
	}
}
