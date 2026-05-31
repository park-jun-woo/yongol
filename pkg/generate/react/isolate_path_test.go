//ff:func feature=gen-react type=test control=sequence
//ff:what resolveOpenapiTsBinary env→PATH→npx 해결 순서·부재 에러 검증
package react

import (
	"testing"
)

func isolatePATH(t *testing.T) {
	t.Helper()
	empty := t.TempDir()
	t.Setenv("PATH", empty)
	t.Setenv("YONGOL_OPENAPI_TS_PROJECT_DIR", "")
	t.Setenv("YONGOL_SWC_PROJECT_DIR", "")
}
