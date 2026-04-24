//ff:type feature=manifest type=util
//ff:what insertIntoRe + onConflictDoNothingRe — INSERT / ON CONFLICT 매처 정규식
package ddl

import (
	"regexp"
)

// insertIntoRe matches the start of a top-level INSERT, capturing the
// target table name. We accept the common forms used by PostgreSQL
// (optional schema prefix is NOT supported in v1 — keep parser simple).
var insertIntoRe = regexp.MustCompile(`(?i)^\s*INSERT\s+INTO\s+([A-Za-z_][A-Za-z0-9_]*)`)

// onConflictDoNothingRe matches `ON CONFLICT ... DO NOTHING` anywhere in
// the collected INSERT body. Whitespace is flexible; the `...` between
// ON CONFLICT and DO NOTHING (target list, action, etc.) is unrestricted.
var onConflictDoNothingRe = regexp.MustCompile(`(?is)ON\s+CONFLICT\b[^;]*\bDO\s+NOTHING\b`)
