//ff:func feature=validate type=test control=sequence topic=migration-safety
//ff:what MIG-003 positive — 지정된 data migration 파일이 없으면 ERROR

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestMIG003_Positive_MissingFile(t *testing.T) {
	diags := Mig003DataMigrationMissing([]string{"migrations_data/doesnotexist.sql"})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d", len(diags))
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR, got %v", diags[0].Level)
	}
}
