//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestDF_01_SubscribeUnmarshal_Guarded — generate_subscribe_method.go 소스가 Unmarshal err 체크 한 줄을 유지하는지 회귀 방지

package ssac

import (
	"os"
	"strings"
	"testing"
)

// TestDF_01_SubscribeUnmarshal_Guarded reads the generator source itself
// and asserts that the body-emitting `json.Unmarshal` literal still includes
// the inline `err != nil` guard (Phase003 DF-01). This is a source-level
// regression probe — if a future refactor drops the guard, every subscribe
// handler would silently swallow malformed queue messages. The check is
// intentionally textual (strings.Contains) to keep the test self-contained
// and free from generator fixture plumbing.
func TestDF_01_SubscribeUnmarshal_Guarded(t *testing.T) {
	src, err := os.ReadFile("generate_subscribe_method.go")
	if err != nil {
		t.Fatalf("read generator source: %v", err)
	}
	if !strings.Contains(string(src), `if err := json.Unmarshal(msg, &message); err != nil { return err }`) {
		t.Fatal("subscribe generator must emit guarded json.Unmarshal (DF-01)")
	}
}
