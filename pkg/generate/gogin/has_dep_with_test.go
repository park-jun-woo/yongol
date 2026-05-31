//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestResolveGoModDeps — tracing off/on(otlp/stdout/default) 의존성 병합 분기 검증
package gogin

import (
	"strings"
)

func hasDepWith(deps []string, substr string) bool {
	for _, d := range deps {
		if strings.Contains(d, substr) {
			return true
		}
	}
	return false
}
