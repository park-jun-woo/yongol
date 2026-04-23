//ff:type feature=migration type=model
//ff:what EmitOptions — EmitSQL 헤더 튜닝 (yongol 버전 + 생성 시각)
package migration

import "time"

// EmitOptions tunes the header yongol writes above each migration.
type EmitOptions struct {
	YongolVersion string    // "v0.1.21"
	GeneratedAt   time.Time // zero => time.Now() at call time
}
