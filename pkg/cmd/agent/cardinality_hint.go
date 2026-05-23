//ff:func feature=agent type=helper control=selection
//ff:what cardinalityHint — op와 path로부터 sqlc cardinality 어노테이션 추론

package agent

import "strings"

// cardinalityHint infers the sqlc cardinality annotation from the operation and path.
func cardinalityHint(op, path string) string {
	// Extract HTTP method from operationId naming convention
	opLower := strings.ToLower(op)

	switch {
	case strings.HasPrefix(opLower, "get") || strings.HasPrefix(opLower, "list"):
		// GET /{id} -> :one, GET / -> :many
		if strings.Contains(path, "{") {
			return ":one"
		}
		return ":many"
	case strings.HasPrefix(opLower, "create") || strings.HasPrefix(opLower, "post"):
		return ":one RETURNING"
	case strings.HasPrefix(opLower, "update") || strings.HasPrefix(opLower, "put"):
		return ":one RETURNING"
	case strings.HasPrefix(opLower, "delete") || strings.HasPrefix(opLower, "remove"):
		return ":exec"
	default:
		return ":one"
	}
}
