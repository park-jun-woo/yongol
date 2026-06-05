//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestClassifyResponseShapesMissingDir — classifyResponseShapes 단위 테스트 (embedded struct vs schema alias 분류)

package ssac

import (
	"path/filepath"
	"testing"
)

func TestClassifyResponseShapesMissingDir(t *testing.T) {
	if got := classifyResponseShapes(filepath.Join(t.TempDir(), "absent")); got != nil {
		t.Errorf("classifyResponseShapes(absent) = %v, want nil", got)
	}
}

// TestClassifyResponseShapesSkips covers the best-effort skip branches:
// non-.go files, sub-directories, unparseable Go, non-type declarations,
// and non-JSONResponse type names are all ignored.
