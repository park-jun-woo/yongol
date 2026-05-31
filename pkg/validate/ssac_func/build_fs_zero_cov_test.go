//ff:func feature=validate type=test control=sequence topic=ssac-func
//ff:what zz_zerocov_test — ssac_func 0% 규칙 (Run / xfs42 / xfs43 / xfs45 / xsf46) 단위 테스트
package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func buildFSZeroCov(seq parsessac.Sequence, spec funcspec.FuncSpec) *yongol.Fullstack {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:      "Handler",
			FileName:  "svc/handler.ssac",
			Sequences: []parsessac.Sequence{seq},
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{spec},
	}
	fs.SetGround(ground.Build(fs))
	return fs
}
