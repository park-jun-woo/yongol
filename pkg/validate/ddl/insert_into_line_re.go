//ff:type feature=validate type=util
//ff:what insertIntoLineRe — validate 내부 INSERT INTO 라인 매처 (parser 와 독립)

package ddl

import (
	"regexp"
)

// insertIntoLineRe captures the target table name from an INSERT INTO line
// (used only for error messages). Duplicated locally to keep validate
// self-contained (parser owns the canonical scanner).
var insertIntoLineRe = regexp.MustCompile(`(?i)^\s*INSERT\s+INTO\s+([A-Za-z_][A-Za-z0-9_]*)`)
