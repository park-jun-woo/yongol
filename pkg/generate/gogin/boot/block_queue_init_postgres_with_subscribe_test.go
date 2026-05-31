//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockQueueInit — queue.Init + SetBackend + Subscribe + Start + defer Close
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
