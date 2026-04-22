//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what parseDDLContent — SQL 텍스트에서 CREATE TABLE 문을 라인 단위로 파싱
package ddl

import "strings"

func parseDDLContent(content string, tables map[string]*Table, file string) {
	lines := strings.Split(content, "\n")
	var currentTable string
	pendingArchived := false
	for i, line := range lines {
		lineNum := i + 1
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if isArchivedAnnotation(trimmed) {
			pendingArchived = true
			continue
		}
		currentTable = parseDDLLine(line, currentTable, tables, pendingArchived, file, lineNum)
		// annotation was consumed by whichever line followed it
		pendingArchived = false
	}
}
