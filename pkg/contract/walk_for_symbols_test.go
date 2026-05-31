//ff:func feature=contract type=test control=sequence
//ff:what test: TestWalkForSymbols — body 2-pass walk 으로 SqlcQueries·CallTargets·DDLFields 정렬 분류 검증
package contract

import (
	"reflect"
	"testing"
)

func TestWalkForSymbols(t *testing.T) {
	fset, body := bodyFromFunc(t,
		"server.Queries.ListItems(ctx)\n"+
			"server.Queries.GetItem(ctx, 1)\n"+
			"billing.Charge(ctx)\n"+
			"notify.Send(ctx)\n"+
			"_ = user.Name\n"+
			"_ = user.Email\n")

	sym := walkForSymbols(fset, body)

	wantQueries := []string{"GetItem", "ListItems"}
	if !reflect.DeepEqual(sym.SqlcQueries, wantQueries) {
		t.Errorf("SqlcQueries: got %v want %v", sym.SqlcQueries, wantQueries)
	}
	wantCalls := []string{"billing.Charge", "notify.Send"}
	if !reflect.DeepEqual(sym.CallTargets, wantCalls) {
		t.Errorf("CallTargets: got %v want %v", sym.CallTargets, wantCalls)
	}
	// `server.Queries` (the receiver of the .ListItems/.GetItem calls) is
	// itself a non-call SelectorExpr with an exported tail, so it is also
	// collected as a DDL-shaped field alongside the user.* accesses.
	wantFields := []string{"server.Queries", "user.Email", "user.Name"}
	if !reflect.DeepEqual(sym.DDLFields, wantFields) {
		t.Errorf("DDLFields: got %v want %v", sym.DDLFields, wantFields)
	}
}
