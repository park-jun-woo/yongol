//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what xss60ModelToTableName — 모델명 → DDL 테이블명 변환 (PascalCase → snake_case → plural)

package ssac

import (
	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"
)

// xss60ModelToTableName mirrors ssac_sqlc.modelToTableName without introducing
// a cross-package dependency.
func xss60ModelToTableName(model string) string {
	return inflection.Plural(strcase.ToSnake(model))
}
