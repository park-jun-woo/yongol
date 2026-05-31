//ff:func feature=generate type=test control=sequence
//ff:what TestRunByName_ZeroCov — runBackend/runFrontend/runMigration/runMigrationStep/runSTMLCodegen 직접 호출
package generate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunMigrationStep_ZeroCov(t *testing.T) {
	// Empty SpecsDir → early-return nil.
	fs := &yongol.Fullstack{SpecsDir: ""}
	if err := runMigrationStep(fs, t.TempDir(), &generateConfig{}); err != nil {
		t.Fatalf("runMigrationStep empty SpecsDir: %v", err)
	}
}
