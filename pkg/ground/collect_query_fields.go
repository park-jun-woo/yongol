//ff:func feature=rule type=util control=iteration dimension=1
//ff:what collectQueryFields — Args/Inputs에서 query.<field> 참조를 추출해 set에 추가
package ground

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func collectQueryFields(args []parseArg, inputs map[string]string, dst rule.StringSet) {
	for _, a := range args {
		if a.Source == "query" && a.Field != "" {
			dst[a.Field] = true
		}
	}
	for _, v := range inputs {
		if strings.HasPrefix(v, "query.") {
			dst[strings.TrimPrefix(v, "query.")] = true
		}
	}
}
