//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteEventBus — FastAPI EventBus stub 파일 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteEventBus(t *testing.T) {
	t.Run("WritesEventBusStub", func(t *testing.T) {
		appDir := t.TempDir()
		if err := writeEventBus(appDir); err != nil {
			t.Fatalf("writeEventBus error: %v", err)
		}
		data, err := os.ReadFile(filepath.Join(appDir, "dependencies", "event_bus.py"))
		if err != nil {
			t.Fatalf("expected event_bus.py: %v", err)
		}
		if !strings.Contains(string(data), "class EventBus") {
			t.Errorf("event_bus.py missing EventBus class")
		}
	})

	t.Run("MkdirFails", func(t *testing.T) {
		base := t.TempDir()
		filePath := filepath.Join(base, "file")
		if err := os.WriteFile(filePath, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeEventBus(filePath); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})
}
