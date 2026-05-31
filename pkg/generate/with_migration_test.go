//ff:func feature=generate type=test control=sequence
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"bytes"
	"testing"
)

func TestWithMigration(t *testing.T) {
	var buf bytes.Buffer
	opt := WithMigration(MigrationHook{Version: "v1", Logger: &buf})
	var c generateConfig
	opt(&c)
	if c.migration.Version != "v1" {
		t.Errorf("WithMigration did not set version, got %q", c.migration.Version)
	}
}
