//ff:type feature=rule type=model
//ff:what parseArg — pkg/parser/ssac.Arg 재노출 (패키지 경계 최소화)
package ground

import parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// parseArg aliases pkg/parser/ssac.Arg for internal use within pkg/ground helpers.
type parseArg = parsessac.Arg
