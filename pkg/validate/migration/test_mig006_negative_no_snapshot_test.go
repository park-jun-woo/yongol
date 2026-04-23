//ff:func feature=validate type=test control=sequence topic=migration-snapshot
//ff:what MIG-006 negative — 스냅샷 파일이 없으면 no-op

package migration

import (
	"testing"
)

func TestMIG006_Negative_NoSnapshot(t *testing.T) {
	tmp := t.TempDir()
	diags := Mig006SnapshotDrift(tmp)
	if len(diags) != 0 {
		t.Errorf("absent snapshot = no diag, got %+v", diags)
	}
}
