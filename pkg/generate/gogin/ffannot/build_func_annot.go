//ff:func feature=gen-gogin type=util control=sequence
//ff:what BuildFuncAnnot — //ff:func feature=... type=... control=... [dimension=N] [topic=...] 문자열 조립

package ffannot

import (
	"fmt"
	"strings"
)

// BuildFuncAnnot renders a single-line //ff:func annotation.
// Produces "//ff:func feature=X type=Y control=Z[ dimension=N][ topic=T]".
// Panics when Feature, Type, or Control is empty — generators must supply all three.
func BuildFuncAnnot(a FuncAnnot) string {
	if a.Feature == "" || a.Type == "" || a.Control == "" {
		panic(fmt.Sprintf("ffannot: empty feature/type/control: %+v", a))
	}
	var sb strings.Builder
	sb.WriteString("//ff:func feature=")
	sb.WriteString(a.Feature)
	sb.WriteString(" type=")
	sb.WriteString(a.Type)
	sb.WriteString(" control=")
	sb.WriteString(a.Control)
	writeDimension(&sb, a)
	writeTopic(&sb, a)
	return sb.String()
}
