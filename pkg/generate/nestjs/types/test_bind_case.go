//ff:type feature=gen-nestjs type=test-helper
//ff:what bindCase — NestJS Bind 매트릭스 테스트 케이스 (family + opts + 기대값)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindCase covers one row in the NestJS type binding matrix: a PGFamily
// plus BindOpts input, and the expected core fields of the resulting
// ir.TypeBinding.
type bindCase struct {
	name           string
	family         typemap.PGFamily
	opts           ir.BindOpts
	wantDBType     string
	wantAPIType    string
	wantToDBExpr   string
	wantToAPIExpr  string
	wantToRespExpr string
	wantNilCheck   string
	wantSupported  bool
}
