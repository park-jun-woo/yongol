//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what DetectControl — 생성된 함수 body 라인에서 지배적 제어구조(sequence/selection/iteration) 추론

package ffannot

// Control constants mirror filefunc's A10~A14 vocabulary.
const (
	ControlSequence  = "sequence"
	ControlSelection = "selection"
	ControlIteration = "iteration"
)

// DetectControl inspects function body lines (already rendered, without enclosing
// braces or signature) and returns the depth-1 control structure that appears
// first in source order — matching filefunc's own parse.DetectControl semantics.
//
// Depth is tracked via balanced { } counts so nested for/switch inside another
// block does not register as depth-1.
//
// The function is intentionally lenient: generators that already know the
// structure should call Build*Annot directly with a literal control value.
// DetectControl exists for templates that assemble body lines dynamically.
func DetectControl(bodyLines []string) string {
	depth := 0
	for _, raw := range bodyLines {
		kind, delta := scanLine(raw, depth)
		if kind != "" {
			return kind
		}
		depth = clampDepth(depth + delta)
	}
	return ControlSequence
}
