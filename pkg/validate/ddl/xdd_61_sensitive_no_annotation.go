//ff:func feature=validate type=rule control=iteration dimension=3 topic=ddl-structural
//ff:what XDD-61 — 민감 패턴 컬럼 @sensitive 없음

package ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdd61SensitiveNoAnnotation validates XDD-61: db/*.sql의 컬럼 이름이
// 민감 패턴(password, token, ssn 등)에 매치되지만 같은 라인에
// `-- @sensitive` (또는 `-- @nosensitive`) 어노테이션이 없으면 WARNING.
// 어노테이션이 없으면 JSON 응답에 노출될 수 있어 경고한다.
func xdd61SensitiveNoAnnotation(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
				match := matchSensitivePattern(colName)
				if match == "" {
					continue
				}
				if hasInlineSensitiveAnnotation(line) {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  f.path,
					Line:  blk.startLine + offset,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelWarning,
					Message: fmt.Sprintf(
						"[XDD-61] 테이블 %q 컬럼 %q — 민감 패턴 %q 에 매치되지만 -- @sensitive 어노테이션이 없습니다",
						blk.tableName, colName, match),
					Advice: "민감 컬럼 라인 끝에 -- @sensitive 또는 -- @nosensitive 어노테이션을 추가하세요 (없으면 JSON 응답에 노출될 수 있음)",
				})
			}
		}
	}
	return diags
}
