//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what modelToTable — 모델 이름을 DDL 테이블 이름(plural snake_case)으로 변환

package ssac_ddl

import (
	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"
)

// modelToTable converts a model name to a table name.
// e.g. "User" → "users", "Reservation" → "reservations"
func modelToTable(model string) string {
	return inflection.Plural(strcase.ToSnake(model))
}
