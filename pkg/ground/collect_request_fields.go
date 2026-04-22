//ff:func feature=rule type=util control=iteration dimension=1
//ff:what collectRequestFields — Args/Inputs에서 request.<field> 참조를 추출해 set에 추가
package ground

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func collectRequestFields(args []parseArg, inputs map[string]string, dst rule.StringSet) {
	for _, a := range args {
		if a.Source == "request" && a.Field != "" {
			dst[a.Field] = true
		}
	}
	for _, v := range inputs {
		if strings.HasPrefix(v, "request.") {
			dst[strings.TrimPrefix(v, "request.")] = true
		}
	}
}
