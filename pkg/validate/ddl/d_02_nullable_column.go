//ff:func feature=validate type=rule control=iteration dimension=3 topic=ddl-structural
//ff:what D-2 — NOT NULL 누락

package ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d02NullableColumn validates D-2: db/*.sql의 컬럼 정의에 NOT NULL이
// 없으면 ERROR. PRIMARY KEY 컬럼은 암묵적으로 NOT NULL이므로 제외.
func d02NullableColumn(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, f := range files {
		blocks := extractTableBlocks(f)
		for _, blk := range blocks {
			lines := strings.Split(blk.body, "\n")
			for offset, line := range lines {
				trimmed := strings.TrimSpace(line)
				if isSkippableDDLLine(trimmed) {
					continue
				}
				m := columnNameRe.FindStringSubmatch(trimmed)
				if m == nil {
					continue
				}
				colName := m[1]
				upper := strings.ToUpper(trimmed)
				if strings.Contains(upper, "PRIMARY KEY") || strings.Contains(upper, "NOT NULL") {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  f.path,
					Line:  blk.startLine + offset,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[D-2] 테이블 %q 컬럼 %q — NOT NULL 이 없습니다",
						blk.tableName, colName),
					Advice: fmt.Sprintf("%s.%s 에 NOT NULL 제약을 추가하거나 명시적 nullable 의도를 -- @nullable 코멘트로 적으세요", blk.tableName, colName),
				})
			}
		}
	}
	return diags
}
