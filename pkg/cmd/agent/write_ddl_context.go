//ff:func feature=agent type=helper control=sequence
//ff:what writeDDLContext — DDL 테이블 컨텍스트 기록

package agent

import (
	"fmt"
	"strings"
)

func writeDDLContext(b *strings.Builder, specsDir, table string) {
	if table == "" {
		return
	}
	ddl := readDDLForTable(specsDir, table)
	if ddl != "" {
		fmt.Fprintf(b, "DDL (%s.sql):\n%s\n\n", table, ddl)
	}
}
