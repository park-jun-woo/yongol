//ff:func feature=agent type=helper control=sequence
//ff:what readDDLForTable — 테이블명으로 DDL 파일 내용 읽기

package agent

import (
	"os"
	"path/filepath"
)

func readDDLForTable(specsDir, tableName string) string {
	if tableName == "" {
		return ""
	}
	data, err := os.ReadFile(filepath.Join(specsDir, "db", tableName+".sql"))
	if err != nil {
		return ""
	}
	return string(data)
}
