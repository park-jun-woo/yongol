//ff:func feature=gen-react type=util control=sequence
//ff:what variant 컴포넌트의 구조분해 매개변수 목록을 생성한다

package react

import "fmt"

// buildDestructParams builds the destructuring parameter list for a variant component.
func buildDestructParams(variantKeys, sizeKeys []string, defaultVariant, defaultSize string) []string {
	var params []string
	if len(variantKeys) > 0 {
		params = append(params, fmt.Sprintf("variant = '%s'", defaultVariant))
	}
	if len(sizeKeys) > 0 {
		params = append(params, fmt.Sprintf("size = '%s'", defaultSize))
	}
	params = append(params, "className", "...props")
	return params
}
