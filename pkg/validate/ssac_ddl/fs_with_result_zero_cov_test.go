//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what zz_zerocov_test — ssac_ddl 0% (Run / xds12ResultNoDDLTable / collectFuncResultDiags) 단위 테스트
package ssac_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithResultZeroCov() *yongol.Fullstack {
	// One @get sequence binding a @result of type "Workflow" with no DDL
	// table → XDS-12 warning.
	return &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:     "GetWorkflow",
			FileName: "svc/get.ssac",
			Sequences: []ssac.Sequence{{
				Type:   "get",
				Line:   2,
				Result: &ssac.Result{Type: "Workflow", Var: "wf"},
			}},
		}},
	}
}
