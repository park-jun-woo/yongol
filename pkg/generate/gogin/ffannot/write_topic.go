//ff:func feature=gen-gogin type=util control=sequence
//ff:what writeTopic — Topic이 비어있지 않으면 " topic=X" 덧붙이기

package ffannot

import "strings"

// writeTopic appends " topic=X" when a.Topic is non-empty.
func writeTopic(sb *strings.Builder, a FuncAnnot) {
	if a.Topic == "" {
		return
	}
	sb.WriteString(" topic=")
	sb.WriteString(a.Topic)
}
