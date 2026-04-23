//ff:func feature=migration type=util control=iteration dimension=1
//ff:what EmitSQL — Operation 리스트를 BEGIN;…COMMIT; 래핑 + 주석 헤더 붙여 단일 SQL 파일 문자열로
package migration

import (
	"fmt"
	"strings"
	"time"
)

// EmitOptions tunes the header yongol writes above each migration.
type EmitOptions struct {
	YongolVersion string    // "v0.1.21"
	GeneratedAt    time.Time // zero => time.Now() at call time
}

// EmitSQL produces the full text of a migration file from ops.
func EmitSQL(ops []Operation, opt EmitOptions) string {
	ts := opt.GeneratedAt
	if ts.IsZero() {
		ts = time.Now().UTC()
	}
	ver := opt.YongolVersion
	if ver == "" {
		ver = "vX.Y.Z"
	}
	b := strings.Builder{}
	// Header
	fmt.Fprintf(&b, "%s%s at %s\n",
		MigrationGeneratedHeaderPrefix, ver, ts.UTC().Format(time.RFC3339))
	b.WriteString("-- Changes:\n")
	for _, op := range ops {
		fmt.Fprintf(&b, "--   * %s\n", op.Description())
	}
	b.WriteByte('\n')
	// Transaction
	b.WriteString("BEGIN;\n\n")
	for _, op := range ops {
		b.WriteString(op.SQL())
		b.WriteString("\n\n")
	}
	b.WriteString("COMMIT;\n")
	return b.String()
}
