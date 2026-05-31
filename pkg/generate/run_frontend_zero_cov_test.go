//ff:func feature=generate type=test control=sequence
//ff:what TestRunByName_ZeroCov — runBackend/runFrontend/runMigration/runMigrationStep/runSTMLCodegen 직접 호출
package generate

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunFrontend_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{}
	// React branch: react.Generate + runSTMLCodegen (no STML → skip) +
	// copyFrontendComponents (empty SpecsDir).
	if err := runFrontend(fs, t.TempDir(), React); err != nil {
		t.Fatalf("runFrontend React: %v", err)
	}
	// Unknown frontend → error.
	if err := runFrontend(fs, t.TempDir(), FrontendType("nope")); err == nil ||
		!strings.Contains(err.Error(), "unknown frontend") {
		t.Errorf("expected unknown frontend error, got %v", err)
	}
}
