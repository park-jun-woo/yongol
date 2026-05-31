//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XSD-55 — 미참조 DDL 테이블 검출 + @func-managed/@archived 면제
package ssac_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func newXsd55Fullstack(flags rule.StringSet, tables ...ddl.Table) *yongol.Fullstack {
	fs := &yongol.Fullstack{DDLTables: tables}
	g := &rule.Ground{Flags: flags}
	fs.SetGround(g)
	return fs
}
