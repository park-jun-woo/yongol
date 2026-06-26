//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what findMatchingSpec — 응답 타입 이름에 대응하는 FuncSpec 검색

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// findMatchingSpec returns the FuncSpec whose pascalCase(Name)+"Response"
// equals the given respName. Returns nil when no match is found.
func findMatchingSpec(respName string, specs []funcspec.FuncSpec) *funcspec.FuncSpec {
	for i := range specs {
		if pascalCase(specs[i].Name)+"Response" == respName {
			return &specs[i]
		}
	}
	return nil
}
