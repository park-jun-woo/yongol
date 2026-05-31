//ff:func feature=gen-filefunc type=test control=iteration dimension=1
//ff:what TestEnsureGenFilefuncEntry — anchor feature 삽입 + 기존 값 보존 검증
package filefunc

import (
	"testing"
)

func TestEnsureGenFilefuncEntry(t *testing.T) {
	dst := map[string]string{"main": "custom main"} // existing value must be preserved
	ensureGenFilefuncEntry(dst)

	if got := dst["main"]; got != "custom main" {
		t.Errorf("existing 'main' should be preserved, got %q", got)
	}
	for name := range anchorFeatures {
		if _, ok := dst[name]; !ok {
			t.Errorf("anchor feature %q not inserted: %v", name, dst)
		}
	}
	if dst["gen-filefunc"] != anchorFeatures["gen-filefunc"] {
		t.Errorf("expected anchor desc for gen-filefunc, got %q", dst["gen-filefunc"])
	}
}
