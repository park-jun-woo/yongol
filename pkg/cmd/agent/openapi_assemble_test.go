//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestAssembleFullOpenAPI — path 블록 정렬·조립, pathToOps 매핑 및 빈 ops fallback offset 검증

package agent

import (
	"strings"
	"testing"
)

func TestAssembleFullOpenAPI(t *testing.T) {
	pathBlocks := map[string]any{
		"/users": map[string]any{
			"get": map[string]any{"operationId": "ListUsers"},
		},
		"/orgs": map[string]any{
			"get": map[string]any{"operationId": "ListOrgs"},
		},
	}
	pathToOps := map[string][]string{
		"/users": {"ListUsers"},
		// "/orgs" intentionally omitted → triggers the empty-ops fallback that
		// uses the path key as the op name.
	}

	out, offsets := assembleFullOpenAPI("MyAPI", pathBlocks, pathToOps)

	if !strings.Contains(out, "title: MyAPI") {
		t.Errorf("missing project title: %q", out)
	}
	if !strings.Contains(out, "openapi: \"3.1.0\"") {
		t.Errorf("missing openapi version header")
	}
	if !strings.Contains(out, "/users") || !strings.Contains(out, "/orgs") {
		t.Errorf("missing path blocks: %q", out)
	}
	if !strings.Contains(out, "bearerAuth") || !strings.Contains(out, "schemas:") {
		t.Errorf("missing components section: %q", out)
	}

	// Paths are sorted: /orgs comes before /users.
	if strings.Index(out, "/orgs") > strings.Index(out, "/users") {
		t.Errorf("paths not sorted (orgs should precede users): %q", out)
	}

	// One offset for ListUsers (from pathToOps) and one for the /orgs fallback.
	byOp := map[string]pathOffset{}
	for _, o := range offsets {
		byOp[o.Op] = o
	}
	if _, ok := byOp["ListUsers"]; !ok {
		t.Errorf("missing ListUsers offset: %+v", offsets)
	}
	// Fallback uses the path key as the op.
	if _, ok := byOp["/orgs"]; !ok {
		t.Errorf("expected /orgs fallback offset: %+v", offsets)
	}
	for _, o := range offsets {
		if o.EndLine < o.StartLine {
			t.Errorf("offset %+v has EndLine < StartLine", o)
		}
	}
}

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
