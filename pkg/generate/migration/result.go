//ff:type feature=migration type=model
//ff:what Result — Generate 의 반환 구조 (마이그레이션 파일/스냅샷/경고)
package migration

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// Result is what Generate returns to the caller.
type Result struct {
	Mode          Mode
	MigrationFile string // relative to artifactsDir, "" on noop
	SnapshotFile  string // relative to specsDir
	OpsCount      int
	Operations    []Operation
	Hints         *Hints
	Warnings      []diagnostic.Diagnostic // non-blocking MIG-004/005 etc
}
