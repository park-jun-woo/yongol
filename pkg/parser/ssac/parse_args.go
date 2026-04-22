//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what 쉼표 분리 인자를 파싱하여 []Arg 반환
package ssac

import "strings"

// parseArgs는 쉼표 분리 인자를 파싱한다.
func parseArgs(s string) []Arg {
	parts := strings.Split(s, ",")
	var args []Arg
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		args = append(args, parseArg(p))
	}
	return args
}
