//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.isPgtypeAlreadyPointer — dotted 필드 접근의 pgtype 변환 결과가 이미 포인터 타입인지 판별

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
)

// isPgtypeAlreadyPointer returns true when the dotted expression (e.g.
// "user.Name") refers to a pgtype column whose GoTypeBinding.ApiField is
// already a pointer type (*string, *time.Time, etc.). In that case the
// pgtypex Ptr-variant bridge returns *T directly, so no additional ptrOf
// wrapping is needed.
func (g *methodGen) isPgtypeAlreadyPointer(varExpr string) bool {
	parts := strings.SplitN(varExpr, ".", 2)
	if len(parts) != 2 {
		return false
	}
	modelName, ok := g.VarTypes[parts[0]]
	if !ok {
		return false
	}
	modelName = strings.TrimPrefix(modelName, "[]")
	// Pass the PascalCase sqlc field name directly. lookupDDLColumn applies
	// caseconv.PascalToSnake internally; a prior re-lowercasing step broke
	// acronym fields (ID→"iD"→"i_d" miss), wrongly returning false and
	// causing optional pgtype fields to be double-pointer wrapped.
	col := lookupDDLColumn(g.DDLTables, modelName, parts[1])
	if col == nil {
		return false
	}
	binding := types.MapPGType(*col)
	return strings.HasPrefix(binding.ApiField, "*")
}
