//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 gogin/ssac.Generate 통합 커버리지
package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	pssac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateSubscribeMethod_Integ(t *testing.T) {
	sf := pssac.ServiceFunc{
		Name:     "OnOrderCompleted",
		FileName: "on_order_completed.ssac",
		Feature:  "webhook",
		Param:    &pssac.ParamInfo{TypeName: "OrderCompletedMessage"},
		Subscribe: &pssac.SubscribeInfo{
			Topic:       "order.completed",
			MessageType: "OrderCompletedMessage",
		},
		Structs: []pssac.StructInfo{{
			Name: "OrderCompletedMessage",
			Fields: []pssac.StructField{
				{Name: "OrderID", Type: "int64"},
				{Name: "Status", Type: "string"},
			},
		}},
		Sequences: []pssac.Sequence{{
			Type:  "publish",
			Topic: "order.audited",
			Inputs: map[string]string{
				"OrderID": "message.OrderID",
			},
		}},
	}

	serviceDir := t.TempDir()
	fs := &yongol.Fullstack{ServiceFuncs: []pssac.ServiceFunc{sf}}
	if err := generateSubscribeMethod(sf, fs, serviceDir, "github.com/example/app", nil); err != nil {
		t.Fatalf("generateSubscribeMethod: %v", err)
	}
	out := filepath.Join(serviceDir, "on_order_completed.go")
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("expected generated subscribe file: %v", err)
	}
	src := string(data)
	for _, want := range []string{
		"func (server *Server) OnOrderCompleted(ctx context.Context, msg []byte) error",
		"json.Unmarshal(msg, &message)",
		"order.completed",
	} {
		if !strings.Contains(src, want) {
			t.Errorf("generated source missing %q\n---\n%s", want, src)
		}
	}
}
