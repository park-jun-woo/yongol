//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what canonicalTableKey — 모델/테이블 공통 정규형(단수 lower-snake) 산출

package ssac_ddl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// canonicalTableKey normalises a model name or a table name to a single
// canonical form (singular lower-snake) so XSD-55 matches the way code
// generators (gogin/ssac, ir) match model↔table. PascalToSnake is idempotent on
// already-snake input, so both "AppConfig" and "app_config" map to "app_config".
func canonicalTableKey(name string) string {
	return caseconv.TableSingular(strings.ToLower(caseconv.PascalToSnake(name)))
}
