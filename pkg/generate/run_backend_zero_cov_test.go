//ff:func feature=generate type=test control=sequence
//ff:what TestRunByName_ZeroCov — runBackend/runFrontend/runMigration/runMigrationStep/runSTMLCodegen 직접 호출
package generate

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunBackend_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{}
	// FastAPI on an empty Fullstack writes the scaffold and returns nil.
	if err := runBackend(fs, t.TempDir(), FastAPI); err != nil {
		t.Fatalf("runBackend FastAPI: %v", err)
	}
	// Unknown backend → error.
	if err := runBackend(fs, t.TempDir(), BackendType("nope")); err == nil ||
		!strings.Contains(err.Error(), "unknown backend") {
		t.Errorf("expected unknown backend error, got %v", err)
	}
}
