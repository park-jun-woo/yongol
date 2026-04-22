//ff:func feature=gen-gogin type=util control=sequence
//ff:what refName — "#/components/schemas/X" 에서 "X" 추출

package ssac

import "strings"

// refName extracts "Workflow" from "#/components/schemas/Workflow".
func refName(ref string) string {
	idx := strings.LastIndex(ref, "/")
	if idx < 0 {
		return ref
	}
	return ref[idx+1:]
}
