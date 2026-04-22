//ff:func feature=manifest type=util control=iteration dimension=1
//ff:what enumToStrings — OpenAPI enum []any를 []string으로 변환
package openapi

import "fmt"

func enumToStrings(enums []any) []string {
	var result []string
	for _, e := range enums {
		result = append(result, fmt.Sprint(e))
	}
	return result
}
