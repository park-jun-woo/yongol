//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what isAuthRequiredTable — manifest auth 가 요구하는 테이블이면 true (e.g. refresh_tokens)

package ssac_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// isAuthRequiredTable reports whether the DDL table is internally required by
// the manifest auth subsystem. When backend.auth is configured, yongol itself
// mandates refresh_tokens (XNA-90). Such tables have no SSaC @model reference
// because they are consumed by the generated auth middleware, not by user SSaC
// service functions. Flagging them as unreferenced (XSD-55) is a false positive.
func isAuthRequiredTable(fs *yongol.Fullstack, table string) bool {
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return false
	}
	// Auth-required tables. Currently only refresh_tokens; if additional
	// auth-internal tables are added, extend this set.
	return table == "refresh_tokens"
}
