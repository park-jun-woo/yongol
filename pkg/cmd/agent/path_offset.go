//ff:type feature=agent type=helper
//ff:what pathOffset — OpenAPI YAML에서 단일 op의 path 블록 라인 범위

package agent

// pathOffset records the line range of a single op's path block in the
// assembled OpenAPI YAML.
type pathOffset struct {
	Op        string
	Path      string // e.g. /workflows/{id}
	StartLine int    // 1-based inclusive
	EndLine   int    // 1-based inclusive
}
