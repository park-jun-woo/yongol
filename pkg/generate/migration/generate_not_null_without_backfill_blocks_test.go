//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestGenerate_NotNullWithoutBackfill_Blocks — NOT NULL + backfill 없으면 MIG-002 로 차단
package migration

import (
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestGenerate_NotNullWithoutBackfill_Blocks(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "t.sql", `CREATE TABLE t (id BIGSERIAL PRIMARY KEY);`)
	if _, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()}); err != nil {
		t.Fatalf("initial: %v", err)
	}
	writeSpec(t, filepath.Join(specsDir, "db"), "t.sql", `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, active BOOLEAN NOT NULL);`)
	_, diags, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()})
	if err == nil {
		t.Fatalf("expected block, got nil")
	}
	foundMIG002 := false
	for _, d := range diags {
		if strings.Contains(d.Message, "MIG-002") {
			foundMIG002 = true
		}
	}
	if !foundMIG002 {
		t.Errorf("expected MIG-002, got %+v", diags)
	}
}
