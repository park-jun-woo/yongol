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

func TestRunMigration_ZeroCov(t *testing.T) {
	// Empty specsDir → migration.Generate runs in noop mode, nil logger path.
	diags, err := runMigration(t.TempDir(), t.TempDir(), MigrationHook{Version: "v0.0.0"})
	if err != nil {
		t.Fatalf("runMigration: %v", err)
	}
	_ = diags
}

func TestRunMigrationStep_ZeroCov(t *testing.T) {
	// Empty SpecsDir → early-return nil.
	fs := &yongol.Fullstack{SpecsDir: ""}
	if err := runMigrationStep(fs, t.TempDir(), &generateConfig{}); err != nil {
		t.Fatalf("runMigrationStep empty SpecsDir: %v", err)
	}
}

func TestRunSTMLCodegen_ZeroCov(t *testing.T) {
	// nil fs → nil.
	if err := runSTMLCodegen(nil, t.TempDir()); err != nil {
		t.Errorf("nil fs should be nil, got %v", err)
	}
	// no STML pages → nil.
	if err := runSTMLCodegen(&yongol.Fullstack{}, t.TempDir()); err != nil {
		t.Errorf("no STML pages should be nil, got %v", err)
	}
}
