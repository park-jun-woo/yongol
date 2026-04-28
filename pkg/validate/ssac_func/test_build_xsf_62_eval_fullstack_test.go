//ff:func feature=validate type=test-helper control=sequence topic=func-check
//ff:what buildXsf62EvalFullstack — TestXsf62EvalOnlyRef 용 최소 Fullstack 조립

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildXsf62EvalFullstack returns the minimal Fullstack used by the BUG-002
// integration test: one ServiceFunc carrying the supplied sequences plus a
// Func Spec billing.isZeroBalance to be checked by XSF-62.
func buildXsf62EvalFullstack(seqs []ssac.Sequence) *yongol.Fullstack {
	return &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:      "Charge",
			FileName:  "service/charge.ssac",
			Sequences: seqs,
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{{
			Package: "billing",
			Name:    "isZeroBalance",
			Line:    10,
		}},
	}
}
