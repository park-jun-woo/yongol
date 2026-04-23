//ff:type feature=generate type=model
//ff:what MigrationHook — DDL 마이그레이션 단계 설정 (버전/로거)
package generate

import "io"

// MigrationHook configures the DDL migration step. Version is embedded
// in emitted headers. Logger receives human-readable status lines (ok to
// pass io.Discard / nil).
type MigrationHook struct {
	Version string
	Logger  io.Writer
}
