//ff:func feature=gen-react-stml type=test control=sequence
//ff:what 단순 stml 헬퍼 (string/zod/param/login) 묶음 커버 — coverage attribution 으로 다수 함수 PASS

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestStringHelpers_ZeroCov(t *testing.T) {
	if formatFloat(3) != "3" || formatFloat(2.5) != "2.5" {
		t.Errorf("formatFloat wrong: %q %q", formatFloat(3), formatFloat(2.5))
	}
	if indentStr(3) != "   " {
		t.Errorf("indentStr wrong: %q", indentStr(3))
	}
	if joinWords([]string{"a", "b"}) != "a b" {
		t.Errorf("joinWords wrong")
	}
	if toUpperFirst("abc") != "Abc" || toUpperFirst("") != "" {
		t.Errorf("toUpperFirst wrong")
	}
	if toLowerFirst("Abc") != "abc" || toLowerFirst("") != "" {
		t.Errorf("toLowerFirst wrong")
	}
	if pascalToLabel("RoomID") == "" || pascalToLabel("") != "" {
		t.Errorf("pascalToLabel wrong: %q", pascalToLabel("RoomID"))
	}
	if snakeToLabel("first_name") == "" {
		t.Errorf("snakeToLabel wrong")
	}
}

func TestOperationPredicates_ZeroCov(t *testing.T) {
	if !isDeleteOperation("DeleteX") || isDeleteOperation("GetX") {
		t.Errorf("isDeleteOperation wrong")
	}
	if !isLoginAction("Login") || isLoginAction("X") {
		t.Errorf("isLoginAction wrong")
	}
	if !hasLoginAction([]stmlparser.ActionBlock{{OperationID: "Login"}}) {
		t.Errorf("hasLoginAction should be true")
	}
	if hasLoginAction([]stmlparser.ActionBlock{{OperationID: "X"}}) {
		t.Errorf("hasLoginAction should be false")
	}
	types := map[string]map[string]string{"GetX": {"id": "integer"}}
	if !isIntegerParam("GetX", "id", types) {
		t.Errorf("isIntegerParam should be true")
	}
	if isIntegerParam("GetX", "id", nil) {
		t.Errorf("isIntegerParam nil map should be false")
	}
	if isIntegerParam("Unknown", "id", types) {
		t.Errorf("isIntegerParam unknown op should be false")
	}
}

func TestParamHelpers_ZeroCov(t *testing.T) {
	routeP := stmlparser.ParamBind{Name: "ID", Source: "route.ID"}
	bodyP := stmlparser.ParamBind{Name: "X", Source: "body.X"}
	if extractRouteParamName(routeP) != "ID" {
		t.Errorf("extractRouteParamName route wrong")
	}
	if extractRouteParamName(bodyP) != "" {
		t.Errorf("extractRouteParamName non-route should be empty")
	}
	names := extractRouteParamNames([]stmlparser.ParamBind{routeP, routeP, bodyP})
	if len(names) != 1 || names[0] != "ID" {
		t.Errorf("extractRouteParamNames wrong: %v", names)
	}
	if paramSourceExpr(routeP) != "ID" || paramSourceExpr(bodyP) != "body.X" {
		t.Errorf("paramSourceExpr wrong")
	}
}

func TestZodChains_ZeroCov(t *testing.T) {
	if zodBaseType("integer") != "z.number().int()" {
		t.Errorf("zodBaseType integer wrong")
	}
	if zodBaseType("number") != "z.number()" {
		t.Errorf("zodBaseType number wrong")
	}
	if zodBaseTypeArray("string") == "" {
		t.Errorf("zodBaseTypeArray empty")
	}
	// plain required string.
	if got := zodChain(oapiparser.FieldConstraint{Type: "string", Required: true}); got == "" {
		t.Errorf("zodChain string empty")
	}
	// optional array.
	if got := zodChain(oapiparser.FieldConstraint{Type: "array", ItemType: "integer"}); got == "" {
		t.Errorf("zodChain array empty")
	}
	// enum chain (required + optional).
	if got := zodEnumChain(oapiparser.FieldConstraint{Enum: []string{"a", "b"}, Required: true}); got == "" {
		t.Errorf("zodEnumChain required empty")
	}
	if got := zodChain(oapiparser.FieldConstraint{Enum: []string{"a"}}); got == "" {
		t.Errorf("zodChain enum-routed empty")
	}
}
