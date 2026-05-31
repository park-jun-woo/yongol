//ff:func feature=agent type=test control=sequence
//ff:what TestAssembleFullOpenAPI — path 블록 정렬·조립, pathToOps 매핑 및 빈 ops fallback offset 검증
package agent

import (
	"strings"
	"testing"
)

func TestAssembleFullOpenAPIEmpty(t *testing.T) {
	// No path blocks: still emits a well-formed skeleton with no offsets.
	out, offsets := assembleFullOpenAPI("Empty", nil, nil)
	if !strings.Contains(out, "paths:") || !strings.Contains(out, "components:") {
		t.Errorf("skeleton missing sections: %q", out)
	}
	if len(offsets) != 0 {
		t.Errorf("expected no offsets, got %+v", offsets)
	}
}
