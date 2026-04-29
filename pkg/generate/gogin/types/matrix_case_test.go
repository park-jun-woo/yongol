//ff:type feature=gen-gogin type=test-helper
//ff:what matrixCase — MapPGType 매트릭스 테스트 케이스 (col + 기대값 4 항)

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// matrixCase covers one row in the audit matrix: a column declaration
// (RawType + nullability + optional DEFAULT / CheckEnum) plus the
// expected core fields of the resulting binding. ConvertExpr / InsertExpr
// / ResponseExpr are checked as substring contains because some bindings
// carry helper-name templates that the test should not over-pin.
type matrixCase struct {
	name           string
	col            ddl.Column
	wantSqlcType   string
	wantApiField   string
	wantKind       BindingKind
	wantNeedsOver  bool
	wantSupported  bool
	wantConvertSub string // substring expected in ConvertExpr (skip on "")
}
