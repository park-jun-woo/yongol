//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-sqlc
//ff:what XQS-73 — 부분 SELECT 쿼리 결과 변수의 필드 참조 시 SELECT 컬럼 존재 검증 (ERROR)

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs73PartialSelectField validates XQS-73: when a SSaC @response (field
// mapping or direct target), @empty, or @exists references a field from a
// query result variable, and the query uses a partial SELECT (not SELECT *),
// verify the referenced field exists in the query's SELECT column list.
func xqs73PartialSelectField(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || len(fs.ServiceFuncs) == 0 || len(fs.SQLcQueries) == 0 {
		return nil
	}
	queryMap := buildQueryBodyMap(fs)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		diags = append(diags, checkXqs73Func(fn, queryMap)...)
	}
	return diags
}
