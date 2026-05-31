//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderQueueModule_ZeroCov — QueueModule 소스 생성
package queue

import (
	"strings"
	"testing"
)

func TestRenderQueueModule_ZeroCov(t *testing.T) {
	out := RenderQueueModule()
	for _, want := range []string{"export class QueueModule", "QueueService", "@Global()"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderQueueModule missing %q", want)
		}
	}
}
