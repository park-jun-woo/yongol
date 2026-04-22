//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what generateResponseAssertions — response schema에서 [Asserts] jsonpath exists 생성
package hurl

import (
	"fmt"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

// generateResponseAssertions builds [Asserts] lines from the operation's success response schema.
func generateResponseAssertions(op *openapi3.Operation) []string {
	schema := getSuccessResponseSchema(op)
	if schema == nil {
		return nil
	}
	names := make([]string, 0, len(schema.Properties))
	for name := range schema.Properties {
		names = append(names, name)
	}
	sort.Strings(names)
	var assertions []string
	for _, name := range names {
		if schema.Properties[name] == nil {
			continue
		}
		assertions = append(assertions, fmt.Sprintf(`jsonpath "$.%s" exists`, name))
	}
	return assertions
}
