package queue

import (
	"strings"
	"testing"
)

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderQueueModule_ZeroCov — QueueModule 소스 생성

func TestRenderQueueModule_ZeroCov(t *testing.T) {
	out := RenderQueueModule()
	for _, want := range []string{"export class QueueModule", "QueueService", "@Global()"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderQueueModule missing %q", want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderQueueService_ZeroCov — QueueService 소스 생성

func TestRenderQueueService_ZeroCov(t *testing.T) {
	out := RenderQueueService()
	for _, want := range []string{"export class QueueService", "async publish"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderQueueService missing %q", want)
		}
	}
}
