//ff:func feature=gen-react type=util control=sequence
//ff:what cn() 함수의 인자 목록을 생성한다

package react

import "fmt"

// buildCNArgs builds the cn() function argument list for a variant component.
func buildCNArgs(base string, variantKeys, sizeKeys []string) []string {
	var args []string
	if base != "" {
		args = append(args, fmt.Sprintf("'%s'", base))
	}
	if len(variantKeys) > 0 {
		args = append(args, "variants[variant]")
	}
	if len(sizeKeys) > 0 {
		args = append(args, "sizes[size]")
	}
	args = append(args, "className")
	return args
}
