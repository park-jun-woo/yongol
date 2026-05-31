//ff:func feature=generate type=test control=sequence
//ff:what TestRunByName_ZeroCov — runBackend/runFrontend/runMigration/runMigrationStep/runSTMLCodegen 직접 호출
package generate

import (
	"testing"
)

func TestRunMigration_ZeroCov(t *testing.T) {
	// Empty specsDir → migration.Generate runs in noop mode, nil logger path.
	diags, err := runMigration(t.TempDir(), t.TempDir(), MigrationHook{Version: "v0.0.0"})
	if err != nil {
		t.Fatalf("runMigration: %v", err)
	}
	_ = diags
}
