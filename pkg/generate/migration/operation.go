//ff:type feature=migration type=model
//ff:what Operation — 마이그레이션 단일 스텝(ALTER/CREATE/DROP) 인터페이스
package migration

// Operation is one emitted migration step (one ALTER / CREATE / DROP).
type Operation interface {
	SQL() string
	Description() string
	// Destructive returns true for operations that can destroy data
	// (DROP TABLE/COLUMN, NOT NULL add without backfill, ...).
	Destructive() bool
	// SafetyLevel returns the classification used by check_safety.go.
	SafetyLevel() SafetyLevel
}
