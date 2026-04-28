//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what XDN-02 — backend.auth.user_table 가 가리키는 DDL 테이블 실재 여부 검증

package manifest_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdn02UserTableExists verifies that the table named by
// backend.auth.user_table is parsed from the DDL set (db/*.sql). Skipped
// when auth is inactive or the field is empty — XDN-01 already reports
// the latter. Only the table existence is checked here; column-level
// matches live in XDN-03 / XDN-04.
func xdn02UserTableExists(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !isAuthActive(fs) {
		return nil
	}
	auth := fs.Manifest.Backend.Auth
	if auth.UserTable == "" {
		return nil
	}
	if findUserTable(fs, auth.UserTable) != nil {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:  "manifest.yaml",
		Line:  auth.UserTableLine,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDN-02] backend.auth.user_table=%q does not match any DDL table",
			auth.UserTable,
		),
		Advice: "Create db/" + auth.UserTable + ".sql with a CREATE TABLE " +
			auth.UserTable + " (...) statement, or change user_table to an existing table.",
	}}
}
