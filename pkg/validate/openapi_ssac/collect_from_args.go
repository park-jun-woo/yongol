//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what collectFromArgs — Args 목록에서 request.<Field> 참조를 fields 집합에 추가

package openapi_ssac

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

func collectFromArgs(fields map[string]bool, args []ssac.Arg) {
	for _, arg := range args {
		if arg.Source == "request" && arg.Field != "" {
			fields[arg.Field] = true
		}
	}
}
