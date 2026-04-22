//ff:func feature=gen-gogin type=util control=sequence
//ff:what BuildTypeAnnot — //ff:type feature=... type=... 문자열 조립

package ffannot

import (
	"fmt"
	"strings"
)

// BuildTypeAnnot renders a single-line //ff:type annotation.
// Produces "//ff:type feature=X type=Y".
// Panics when Feature or Type is empty.
func BuildTypeAnnot(a TypeAnnot) string {
	if a.Feature == "" || a.Type == "" {
		panic(fmt.Sprintf("ffannot: empty feature/type: %+v", a))
	}
	var sb strings.Builder
	sb.WriteString("//ff:type feature=")
	sb.WriteString(a.Feature)
	sb.WriteString(" type=")
	sb.WriteString(a.Type)
	return sb.String()
}
