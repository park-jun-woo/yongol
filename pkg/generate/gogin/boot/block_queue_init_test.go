//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockQueueInit — queue.Init + SetBackend + Subscribe + Start + defer Close

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBlockQueueInit_MemoryNoSubscribe(t *testing.T) {
	block := blockQueueInit(prepared.Queue{Backend: "memory"}, nil, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `queue.Init(ctx, "memory")`) {
		t.Errorf("must init memory queue, got:\n%s", body)
	}
	if strings.Contains(body, "SetBackend") {
		t.Errorf("memory backend must not call SetBackend, got:\n%s", body)
	}
	if !strings.Contains(body, "defer queue.Close()") {
		t.Errorf("must defer queue.Close(), got:\n%s", body)
	}
}

func TestBlockQueueInit_PostgresWithSubscribe(t *testing.T) {
	funcs := []ssac.ServiceFunc{
		{Name: "OnOrderCompleted", Subscribe: &ssac.SubscribeInfo{Topic: "order.completed"}},
	}
	block := blockQueueInit(prepared.Queue{Backend: "postgres"}, funcs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "queue.SetBackend(infraqueue.NewPostgres(queries))") {
		t.Errorf("postgres backend must wire infra adapter, got:\n%s", body)
	}
	if !strings.Contains(body, `queue.Subscribe("order.completed", srv.OnOrderCompleted)`) {
		t.Errorf("must register subscriber, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), `infraqueue "example.com/zenflow/internal/infra/queue"`) {
		t.Errorf("postgres backend must import infra queue, got:\n%v", block.Imports)
	}
}
