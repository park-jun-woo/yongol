//ff:func feature=gen-gogin type=util control=sequence
//ff:what writeDimension — control=iteration일 때만 " dimension=N" 덧붙이기 (기본 1)

package ffannot

import (
	"fmt"
	"strings"
)

// writeDimension appends " dimension=N" when a.Control == "iteration".
// Zero/negative Dimension defaults to 1 so every iteration annotation satisfies
// filefunc A15 (dimension required).
func writeDimension(sb *strings.Builder, a FuncAnnot) {
	if a.Control != "iteration" {
		return
	}
	dim := a.Dimension
	if dim <= 0 {
		dim = 1
	}
	fmt.Fprintf(sb, " dimension=%d", dim)
}
