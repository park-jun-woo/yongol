//ff:type feature=gen-gogin type=model
//ff:what FuncAnnot — //ff:func 어노테이션 렌더링 입력 (feature/type/control/dimension/topic)

package ffannot

// FuncAnnot collects the values needed to render a //ff:func annotation line.
// Zero values for optional fields (Topic, Dimension) suppress those parts.
type FuncAnnot struct {
	Feature   string // e.g. "service", "middleware"
	Type      string // e.g. "handler", "generator"
	Control   string // "sequence", "selection", "iteration"
	Dimension int    // required when Control == "iteration"
	Topic     string // optional; e.g. "transaction-boundary"
}
