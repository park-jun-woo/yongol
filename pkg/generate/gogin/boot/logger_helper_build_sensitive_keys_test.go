//ff:func feature=gen-gogin type=test control=sequence
//ff:what loggerHelperBuildSensitiveKeys — redact.DefaultKeys + @sensitive 컬럼 병합 map 생성 헬퍼 소스 반환

package boot

import (
	"strings"
	"testing"
)

func TestLoggerHelperBuildSensitiveKeys(t *testing.T) {
	src := loggerHelperBuildSensitiveKeys()
	for _, must := range []string{
		"func buildSensitiveKeys(extras []string) map[string]bool {",
		"redact.DefaultKeys",
		"for _, k := range extras {",
		"out[k] = true",
		"return out",
	} {
		if !strings.Contains(src, must) {
			t.Errorf("buildSensitiveKeys helper missing %q, got:\n%s", must, src)
		}
	}
}
