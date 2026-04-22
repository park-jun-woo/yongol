//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what ParseTables — db/ 디렉토리의 .sql 파일에서 Table 목록 추출
package ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseTables reads all .sql files in dir and returns parsed tables.
func ParseTables(dir string) ([]Table, []diagnostic.Diagnostic) {
	tables := make(map[string]*Table)
	diags := walkSQLFiles(dir, func(path string, data []byte) []diagnostic.Diagnostic {
		parseDDLContent(string(data), tables, path)
		return nil
	})
	var result []Table
	for _, t := range tables {
		result = append(result, *t)
	}
	return result, diags
}
