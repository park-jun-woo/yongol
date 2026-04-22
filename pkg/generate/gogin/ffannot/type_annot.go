//ff:type feature=gen-gogin type=model
//ff:what TypeAnnot — //ff:type 어노테이션 렌더링 입력 (feature/type)

package ffannot

// TypeAnnot collects the values needed to render a //ff:type annotation line.
type TypeAnnot struct {
	Feature string
	Type    string
}
