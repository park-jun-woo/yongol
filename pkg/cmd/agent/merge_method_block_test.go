//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestMergeMethodBlock — 신규 path 추가 및 기존 path에 method 병합 검증

package agent

import "testing"

func TestMergeMethodBlock(t *testing.T) {
	// New path key: block stored as-is.
	pathBlocks := map[string]any{}
	mergeMethodBlock(pathBlocks, "/users", map[string]any{"get": "g"})
	got, ok := pathBlocks["/users"].(map[string]any)
	if !ok || got["get"] != "g" {
		t.Fatalf("new key = %v, want map with get", pathBlocks["/users"])
	}

	// Existing map: method merged in, keeping prior methods.
	mergeMethodBlock(pathBlocks, "/users", map[string]any{"post": "p"})
	got = pathBlocks["/users"].(map[string]any)
	if got["get"] != "g" || got["post"] != "p" {
		t.Errorf("merged = %v, want both get and post", got)
	}

	// Existing non-map value gets overwritten by block.
	pathBlocks["/x"] = "scalar"
	mergeMethodBlock(pathBlocks, "/x", map[string]any{"get": "g2"})
	if m, ok := pathBlocks["/x"].(map[string]any); !ok || m["get"] != "g2" {
		t.Errorf("overwrite scalar = %v", pathBlocks["/x"])
	}
}
