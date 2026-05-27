//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeQueueModule — QueueModule + QueueService 파일 기록

package nestjs

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/queue"
)

// writeQueueModule writes the QueueModule and QueueService files for DI.
func writeQueueModule(srcDir string) error {
	queueDir := filepath.Join(srcDir, "queue")
	if err := os.MkdirAll(queueDir, 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(queueDir, "queue.service.ts"), []byte(queue.RenderQueueService()), 0o644); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(queueDir, "queue.module.ts"), []byte(queue.RenderQueueModule()), 0o644)
}
