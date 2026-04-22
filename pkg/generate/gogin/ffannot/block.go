//ff:type feature=gen-gogin type=model
//ff:what Block — //ff:func / //ff:type / //ff:what 렌더링에 필요한 입력 집합

package ffannot

// Block holds the annotation header plus optional //ff:what.
// Zero values produce an empty block (no output).
//
// Both Func and Type may be non-zero — EmitAnnotationBlock emits both lines in
// that order. This supports files that contain both a func and a top-level type
// (e.g. IssueTokenRequest struct + IssueToken func) which filefunc A1 requires
// to carry both //ff:func and //ff:type.
//
// All zero → EmitAnnotationBlock returns "".
type Block struct {
	Func FuncAnnot
	Type TypeAnnot
	What string
}
