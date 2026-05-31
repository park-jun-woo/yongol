//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderQueueModule_ZeroCov — QueueModule 소스 생성
package queue

import (
	"strings"
	"testing"
)

func TestRenderQueueService_ZeroCov(t *testing.T) {
	out := RenderQueueService()
	for _, want := range []string{"export class QueueService", "async publish"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderQueueService missing %q", want)
		}
	}
}
