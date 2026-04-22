//ff:func feature=gen-ffhash type=util control=sequence
//ff:what rewriteOne — 파일 1개를 읽어 //ff:checked 를 반영하고 변동이 있을 때만 다시 기록

package ffhash

import (
	"bytes"
	"log/slog"
	"os"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

// rewriteOne reads path, passes the bytes through InjectCheckedLine and
// writes the result back only when it differs from the original. The
// no-diff short-circuit preserves mtime so incremental tooling that
// keys off modification time (e.g. Make, air, goimports caches) sees a
// stable file set across repeated `yongol generate` runs.
//
// Preserved files (DetectPreserved → StatePreserved) are left untouched
// because re-injecting a fresh hash would erase the user's drift signal
// and let the next `yongol generate` overwrite the body.
func rewriteOne(path string) error {
	src, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if state, derr := contract.DetectPreserved(path); derr == nil && state == contract.StatePreserved {
		slog.Info("skipping preserved file during hash refresh", "path", path)
		return nil
	}
	out := InjectCheckedLine(src)
	if bytes.Equal(src, out) {
		return nil
	}
	return os.WriteFile(path, out, 0o644)
}
