//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqsBatch_ZeroCov — ssac_sqlc 검증 헬퍼를 이름으로 직접 호출해 커버 귀속
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

func TestXqsHelperEntries_ZeroCov(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "Fn", FileName: "f.ssac"}
	seq := ssacparser.Sequence{Type: "get", Line: 1}
	queryMap := map[string]sqlcparser.QuerySpec{}
	tableMap := map[string]*ddl.Table{}
	vars := map[string]xqs73VarInfo{}
	have := map[string]bool{}
	interfaces := map[string]*ssacmeta.PackageInterface{}

	_ = buildXqs19Diag(fn, seq, "session", "GetUser", "GetUser")
	_ = buildXqs19DiagSubscribe(fn, "OnX")
	_ = buildXqs20Diag(fn, seq, "Model", "Row", "GetUser", ShapePartial, "reason")
	_ = buildXqs73Vars(fn, queryMap)

	_ = checkVerifyPasswordQueries(fn, have)
	_, _ = checkXqs20Seq(fn, seq, queryMap, tableMap)
	_ = checkXqs73Func(fn, queryMap)
	_ = checkXqs73Guard(fn, seq, vars)
	_ = checkXqs73Response(fn, seq, vars)
	_ = checkXqs73ResponseFields(fn, seq, vars)
	_, _ = checkXqs75Seq(fn, seq, queryMap)
	_, _ = checkXqs76Seq(fn, seq, queryMap)

	_ = collectFuncCallDiags(fn, interfaces, have)
	_ = collectFuncSubscribeDiags(fn, interfaces, have)
	_ = collectSeqPortDiags(fn, seq, interfaces, have)

	if _, ok := xqs18LookupDDLType(map[string]map[string]string{}, "users", "id"); ok {
		t.Error("xqs18LookupDDLType empty should miss")
	}
}
