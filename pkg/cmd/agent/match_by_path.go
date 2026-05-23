//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what matchByPath — 에러 메시지에서 path 키로 op 매핑

package agent

import "strings"

func matchByPath(msg string, offsets []pathOffset) []string {
	var ops []string
	for _, off := range offsets {
		if off.Path != "" && strings.Contains(msg, off.Path) {
			ops = append(ops, off.Op)
		}
	}
	return ops
}
