//ff:type feature=validate type=model topic=states
//ff:what xsm27Case — XSM-27 테이블 기반 테스트 row 타입

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xsm27Case captures one table-driven fixture for TestXsm27StateIntentDeclaration.
type xsm27Case struct {
	name           string
	method         string // "GET" / "POST" / "PUT" / "DELETE"
	path           string // e.g. "/workflows/{id}/execute"
	opID           string
	sequences      []ssac.Sequence
	stateNeutral   bool
	withDiagram    bool
	withDefault    bool
	nonStatefulRes bool // when true skip diagram/default → path resource not stateful
	wantFire       bool
}
