//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what xss11FS -- XSS-11 테스트용 Fullstack 생성 헬퍼

package ssac

import (
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func xss11FS(seqType, model string, result *parsessac.Result, pkg string) *yongol.Fullstack {
	return &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "TestOp",
			FileName: "service/test.ssac",
			Sequences: []parsessac.Sequence{{
				Type:    seqType,
				Model:   model,
				Result:  result,
				Package: pkg,
				Line:    5,
			}},
		}},
	}
}
