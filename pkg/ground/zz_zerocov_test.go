//ff:func feature=rule type=test
//ff:what zz_zerocov_test — ground.registerFuncSpec 0% 커버리지 단위 테스트
package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestRegisterFuncSpec_ZeroCov(t *testing.T) {
	g := newGround()
	sp := &funcspec.FuncSpec{
		Name: "hashPassword",
		RequestFields: []funcspec.Field{
			{Name: "password", Type: "string"},
			{Name: "cost", Type: "int"},
		},
		ResponseFields: []funcspec.Field{
			{Name: "hash", Type: "string"},
		},
	}
	registerFuncSpec(g, sp)

	// Request field types under both Func.request and Struct keys.
	if g.Types["Func.request.hashPassword.password"] != "string" {
		t.Errorf("Func.request type missing: %v", g.Types)
	}
	if g.Types["Struct.HashPasswordRequest.cost"] != "int" {
		t.Errorf("Struct request type missing: %v", g.Types)
	}
	// Schema = ordered request field names.
	schema := g.Schemas["Func.request.hashPassword"]
	if len(schema) != 2 || schema[0] != "password" || schema[1] != "cost" {
		t.Errorf("schema = %v", schema)
	}
	// Response field type.
	if g.Types["Struct.HashPasswordResponse.hash"] != "string" {
		t.Errorf("Struct response type missing: %v", g.Types)
	}
}
