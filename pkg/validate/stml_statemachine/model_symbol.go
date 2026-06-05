//ff:func feature=validate type=helper control=sequence dimension=1 topic=stml-statemachine
//ff:what modelSymbol — 가드 model 접두어를 StateDiagram.Symbol과 동일한 PascalCase로 정규화

package stml_statemachine

import "github.com/ettle/strcase"

// modelSymbol normalizes a guard model prefix (e.g. "workflow") to the same
// PascalCase key used by StateDiagram.Symbol (e.g. "Workflow"), so case-only
// differences between the STML guard and the diagram filename match.
func modelSymbol(model string) string {
	return strcase.ToPascal(model)
}
