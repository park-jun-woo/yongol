//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-39 — @subscribe message type 의 struct 정의 필수 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS39SubscribeMessageRequired(t *testing.T) {
	sub := &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}
	param := &ssac.ParamInfo{TypeName: "OnMsg", VarName: "message"}
	t.Run("Fires_missing_struct", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
			{FileName: "w.ssac", Line: 1, Subscribe: sub, Param: param, Structs: nil},
		}}
		assertDiag(t, s39SubscribeMessageRequired(fs), "[S-39]")
	})
	t.Run("Passes_struct_defined", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
			{FileName: "w.ssac", Line: 1, Subscribe: sub, Param: param, Structs: []ssac.StructInfo{{Name: "OnMsg"}}},
		}}
		assertNoDiag(t, s39SubscribeMessageRequired(fs), "[S-39]")
	})
	t.Run("Skips_http", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "o.ssac", Line: 1, Subscribe: nil}}}
		assertNoDiag(t, s39SubscribeMessageRequired(fs), "[S-39]")
	})
	t.Run("Skips_nil_param", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
			{FileName: "w.ssac", Line: 1, Subscribe: sub, Param: nil},
		}}
		assertNoDiag(t, s39SubscribeMessageRequired(fs), "[S-39]")
	})
	t.Run("Skips_empty_type", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
			{FileName: "w.ssac", Line: 1, Subscribe: sub, Param: &ssac.ParamInfo{TypeName: "", VarName: "message"}},
		}}
		assertNoDiag(t, s39SubscribeMessageRequired(fs), "[S-39]")
	})
	t.Run("Empty", func(t *testing.T) {
		assertNoDiag(t, s39SubscribeMessageRequired(&yongol.Fullstack{}), "[S-39]")
	})
}
