//ff:func feature=gen-gogin type=util control=sequence topic=response
//ff:what isAlreadyPointerApiField — DDL 컬럼의 ApiField 가 이미 *T 포인터인지 판별

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// isAlreadyPointerApiField returns true when the DDL column's mapped
// ApiField already uses a pointer type (*T), meaning ptrOf wrapping
// should be skipped.
func isAlreadyPointerApiField(col *ddl.Column) bool {
	if col == nil {
		return false
	}
	binding := types.MapPGType(*col)
	return strings.HasPrefix(binding.ApiField, "*")
}
