//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what EnsureUnique — 이미 사용된 파일명 집합에 충돌하지 않는 이름을 suffix 로 보장

package fffile

import (
	"fmt"
	"strings"
)

// EnsureUnique returns a file name that does not yet appear in used. When the
// candidate is already taken, numeric suffixes (_2, _3, …) are appended to the
// base (before the ".go" extension) until a free name is found. The chosen
// name is added to used so repeated calls converge on a unique sequence.
//
// A nil used map is treated as empty but cannot be mutated; callers that want
// persistence must pass a non-nil map.
//
// Example:
//
//	used := map[string]bool{"convert_workflow.go": true}
//	EnsureUnique("convert_workflow.go", used) // "convert_workflow_2.go"
func EnsureUnique(candidate string, used map[string]bool) string {
	if candidate == "" {
		return ""
	}
	if used == nil {
		return candidate
	}
	if !used[candidate] {
		used[candidate] = true
		return candidate
	}
	base := strings.TrimSuffix(candidate, ".go")
	for i := 2; ; i++ {
		next := fmt.Sprintf("%s_%d.go", base, i)
		if !used[next] {
			used[next] = true
			return next
		}
	}
}
