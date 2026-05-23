//ff:func feature=gen-gogin type=util control=selection
//ff:what methodGen.buildResponseConvert — @response target의 convert 헬퍼 호출 생성

package ssac

import (
	"fmt"
	"strings"
)

// buildResponseConvert emits the convert + return lines for a @response target
// when the variable has a known model type.
func (g *methodGen) buildResponseConvert(model, target string) []string {
	if elem, ok := strings.CutPrefix(model, "[]"); ok {
		return []string{
			fmt.Sprintf("converted, err := convert%sList(%s)", elem, target),
			"if err != nil { return nil, err }",
			fmt.Sprintf("return api.%s%dJSONResponse(converted), nil",
				g.FuncName, g.SuccessStatus),
		}
	}
	return []string{
		fmt.Sprintf("converted, err := convert%s(%s)", model, target),
		"if err != nil { return nil, err }",
		fmt.Sprintf("return api.%s%dJSONResponse(*converted), nil",
			g.FuncName, g.SuccessStatus),
	}
}
