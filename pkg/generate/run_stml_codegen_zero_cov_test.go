//ff:func feature=generate type=test control=sequence
//ff:what TestRunByName_ZeroCov — runBackend/runFrontend/runMigration/runMigrationStep/runSTMLCodegen 직접 호출
package generate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunSTMLCodegen_ZeroCov(t *testing.T) {
	// nil fs → nil.
	if err := runSTMLCodegen(nil, "", t.TempDir()); err != nil {
		t.Errorf("nil fs should be nil, got %v", err)
	}
	// no STML pages → nil.
	if err := runSTMLCodegen(&yongol.Fullstack{}, "", t.TempDir()); err != nil {
		t.Errorf("no STML pages should be nil, got %v", err)
	}
}
