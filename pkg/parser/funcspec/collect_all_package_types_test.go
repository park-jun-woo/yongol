//ff:func feature=funcspec type=test control=sequence
//ff:what TestCollectAllPackageTypes — rootDir 하위 재귀 struct 수집 검증

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollectAllPackageTypes(t *testing.T) {
	chatSrc := "package chat\n\ntype ChatMessage struct {\n\tTurnID  string `json:\"turn_id\"`\n\tContent string `json:\"content\"`\n}\n\ntype ChatPagination struct {\n\tNextCursor string `json:\"next_cursor\"`\n\tHasMore    bool   `json:\"has_more\"`\n}\n"
	authSrc := "package auth\n\ntype UserInfo struct {\n\tUserID string `json:\"user_id\"`\n\tEmail  string `json:\"email\"`\n}\n"

	t.Run("Basic", func(t *testing.T) {
		root := t.TempDir()
		os.MkdirAll(filepath.Join(root, "chat"), 0o755)
		os.MkdirAll(filepath.Join(root, "auth"), 0o755)
		os.WriteFile(filepath.Join(root, "chat", "types.go"), []byte(chatSrc), 0o644)
		os.WriteFile(filepath.Join(root, "auth", "types.go"), []byte(authSrc), 0o644)

		result, diags := CollectAllPackageTypes(root)
		if len(diags) > 0 {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if _, ok := result["chat"]["ChatMessage"]; !ok {
			t.Error("expected ChatMessage in chat")
		}
		if _, ok := result["chat"]["ChatPagination"]; !ok {
			t.Error("expected ChatPagination in chat")
		}
		if _, ok := result["auth"]["UserInfo"]; !ok {
			t.Error("expected UserInfo in auth")
		}
	})

	t.Run("EmptyDir", func(t *testing.T) {
		result, diags := CollectAllPackageTypes(t.TempDir())
		if len(diags) > 0 {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if len(result) != 0 {
			t.Fatalf("expected empty result, got %d packages", len(result))
		}
	})

	t.Run("FieldResolution", func(t *testing.T) {
		root := t.TempDir()
		os.MkdirAll(filepath.Join(root, "chat"), 0o755)
		os.WriteFile(filepath.Join(root, "chat", "types.go"), []byte(chatSrc), 0o644)
		result, _ := CollectAllPackageTypes(root)
		found := false
		for _, f := range result["chat"]["ChatMessage"] {
			if f.Name == "TurnID" && f.JSONName == "turn_id" {
				found = true
			}
		}
		if !found {
			t.Errorf("expected field TurnID with json tag turn_id, got: %+v", result["chat"]["ChatMessage"])
		}
	})
}
