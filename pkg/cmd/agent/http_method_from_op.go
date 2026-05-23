//ff:func feature=agent type=helper control=selection
//ff:what httpMethodFromOp — operationId에서 HTTP method 추론

package agent

import "strings"

func httpMethodFromOp(op string) string {
	lower := strings.ToLower(op)
	switch {
	case strings.HasPrefix(lower, "list"):
		return "get"
	case strings.HasPrefix(lower, "get"):
		return "get"
	case strings.HasPrefix(lower, "create"):
		return "post"
	case strings.HasPrefix(lower, "update"):
		return "put"
	case strings.HasPrefix(lower, "delete"), strings.HasPrefix(lower, "remove"):
		return "delete"
	default:
		return "post"
	}
}
