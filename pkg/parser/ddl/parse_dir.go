//ff:func feature=ddl type=parser control=sequence
//ff:what 디렉토리 내 모든 .sql 파일을 pg_query_go로 파싱
package ddl

import (
	pg_query "github.com/pganalyze/pg_query_go/v5"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseDir parses all .sql files in the given directory using pg_query_go.
func ParseDir(dir string) ([]*pg_query.ParseResult, []diagnostic.Diagnostic) {
	var results []*pg_query.ParseResult
	diags := walkSQLFiles(dir, func(path string, data []byte) []diagnostic.Diagnostic {
		result, err := pg_query.Parse(string(data))
		if err != nil {
			return []diagnostic.Diagnostic{{
				File:    path,
				Line:    0,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "SQL parse error: " + err.Error(),
			}}
		}
		results = append(results, result)
		return nil
	})
	return results, diags
}
