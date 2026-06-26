//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitFuncResponseConvertListFile — Func Response convert list 파일 emit 검증

package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEmitFuncResponseConvertListFile(t *testing.T) {
	info := funcRespInfo{PkgAlias: "chat", ImportPath: "example.com/app/internal/chat"}

	t.Run("Basic", func(t *testing.T) {
		dir := t.TempDir()
		if err := emitFuncResponseConvertListFile(dir, "example.com/app", "ChatMessage", info, map[string]bool{}, domainGen{}); err != nil {
			t.Fatalf("emitFuncResponseConvertListFile: %v", err)
		}
		entries, _ := os.ReadDir(dir)
		if len(entries) == 0 {
			t.Fatal("expected a file emitted")
		}
		b, _ := os.ReadFile(filepath.Join(dir, entries[0].Name()))
		content := string(b)
		for _, sub := range []string{"convertChatMessageList", "chat.ChatMessage", "[]api.ChatMessage", "convertChatMessage(row)", "example.com/app/internal/chat"} {
			if !strings.Contains(content, sub) {
				t.Errorf("expected %q in output", sub)
			}
		}
	})

	t.Run("DomainGen", func(t *testing.T) {
		dir := t.TempDir()
		dg := domainGen{FuncPrefix: "Admin"}
		if err := emitFuncResponseConvertListFile(dir, "example.com/app", "ChatMessage", info, map[string]bool{}, dg); err != nil {
			t.Fatalf("emitFuncResponseConvertListFile: %v", err)
		}
		entries, _ := os.ReadDir(dir)
		b, _ := os.ReadFile(filepath.Join(dir, entries[0].Name()))
		if !strings.Contains(string(b), "convertAdminChatMessageList") {
			t.Errorf("expected convertAdminChatMessageList in output, got:\n%s", string(b))
		}
	})
}
