//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestDependencies_ZeroCov — 런타임 의존성 맵
package scaffold

import (
	"testing"
)

func TestNestDevDependencies_ZeroCov(t *testing.T) {
	deps := nestDevDependencies()
	if deps["@nestjs/cli"] == "" {
		t.Error("expected @nestjs/cli dev dependency")
	}
}
