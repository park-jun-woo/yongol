//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what modelToTable — OpenAPI 스키마/모델명을 DDL 테이블명으로 변환 (snake + plural)

package openapi_ddl

import (
	"github.com/jinzhu/inflection"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// modelToTable converts a model name (e.g. "User") to a DDL table name (e.g. "users").
func modelToTable(model string) string {
	return inflection.Plural(caseconv.PascalToSnake(model))
}
