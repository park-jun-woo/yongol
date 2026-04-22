//ff:func feature=validate type=rule control=iteration dimension=3 topic=ddl-structural
//ff:what D-3 — FK DEFAULT 0 센티널 레코드 누락

package ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d03SentinelMissing validates D-3: 컬럼이 FK + DEFAULT 0이라면, 참조
// 대상 테이블에 id=0 센티널 레코드(INSERT INTO ... VALUES (0, ...))가
// 존재해야 한다. 없으면 ERROR.
func d03SentinelMissing(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}

	tableContents := allTableContents(files)

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
				if !strings.Contains(upper, "DEFAULT 0") || !strings.Contains(upper, "REFERENCES") {
					continue
				}
				refMatch := referencesRe.FindStringSubmatch(trimmed)
				if refMatch == nil {
					continue
				}
				refTable := refMatch[1]
				refContent, ok := tableContents[refTable]
				if !ok {
					continue
				}
				if hasSentinelInsert(refContent, refTable) {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  f.path,
					Line:  blk.startLine + offset,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[D-3] 테이블 %q 컬럼 %q — FK + DEFAULT 0 이지만 참조 대상 %q 에 id=0 센티널 레코드가 없습니다",
						blk.tableName, colName, refTable),
					Advice: fmt.Sprintf("INSERT INTO %s (id, ...) VALUES (0, ...) ON CONFLICT DO NOTHING; 을 추가하세요", refTable),
				})
			}
		}
	}
	return diags
}
