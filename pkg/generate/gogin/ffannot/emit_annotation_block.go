//ff:func feature=gen-gogin type=util control=selection
//ff:what EmitAnnotationBlock — //ff:func / //ff:type / //ff:what 블록 문자열 생성 (package 선언 직전용)

package ffannot

import "strings"

// EmitAnnotationBlock returns the annotation header as a single string ending
// with a trailing newline. The caller writes it immediately before the
// "package <name>" line.
//
// When both b.Func and b.Type are populated, both lines are emitted (//ff:func
// then //ff:type). When only one is set, only that line appears. When both are
// zero, the result is "" so the caller can opt out cleanly.
//
// Usage:
//
//	header := ffannot.EmitAnnotationBlock(ffannot.Block{
//	    Func: ffannot.FuncAnnot{Feature: "service", Type: "handler", Control: "sequence"},
//	    What: "ActivateWorkflow — activate a workflow",
//	})
//	source := header + "package service\n\n" + bodyStr
func EmitAnnotationBlock(b Block) string {
	var sb strings.Builder
	switch {
	case b.Func.Feature != "" && b.Type.Feature != "":
		sb.WriteString(BuildFuncAnnot(b.Func))
		sb.WriteString("\n")
		sb.WriteString(BuildTypeAnnot(b.Type))
		sb.WriteString("\n")
	case b.Func.Feature != "":
		sb.WriteString(BuildFuncAnnot(b.Func))
		sb.WriteString("\n")
	case b.Type.Feature != "":
		sb.WriteString(BuildTypeAnnot(b.Type))
		sb.WriteString("\n")
	default:
		return ""
	}
	if w := BuildWhat(b.What); w != "" {
		sb.WriteString(w)
		sb.WriteString("\n")
	}
	return sb.String()
}
