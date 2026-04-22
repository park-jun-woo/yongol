//ff:func feature=gen-gogin type=generator control=sequence
//ff:what renderHelperSource — 단일 helperSpec 을 완전한 Go source (annotation + package + body) 로 조립

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// renderHelperSource builds the full Go source for one pointer/deref
// helper: the //ff:func + //ff:what annotation block, the package clause,
// and the pre-formatted function body from spec.code.
func renderHelperSource(spec helperSpec) string {
	var sb strings.Builder
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature: "service",
			Type:    "util",
			Control: "sequence",
			Topic:   "pointer-helper",
		},
		What: spec.what,
	}))
	sb.WriteString("package service\n\n")
	sb.WriteString(spec.code)
	return sb.String()
}
