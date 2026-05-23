//ff:func feature=agent type=helper control=sequence
//ff:what writeSQLcQueryContext — sqlc 쿼리 이름 컨텍스트 기록

package agent

import (
	"fmt"
	"strings"
)

func writeSQLcQueryContext(b *strings.Builder, specsDir, table string) {
	names := readSQLcQueryNames(specsDir, table)
	if len(names) > 0 {
		b.WriteString("sqlc queries:\n")
		for _, n := range names {
			fmt.Fprintf(b, "  %s\n", n)
		}
		b.WriteByte('\n')
	}
}
