//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what xss60FS — XSS-60 테스트용 최소 Fullstack 생성 헬퍼

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xss60FS builds a minimal Fullstack with the given funcs and DDL tables.
func xss60FS(funcs []parsessac.ServiceFunc, tables []ddl.Table) *yongol.Fullstack {
	return &yongol.Fullstack{
		ServiceFuncs: funcs,
		DDLTables:    tables,
	}
}
