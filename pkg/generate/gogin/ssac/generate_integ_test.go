//ff:func feature=gen-gogin type=test control=sequence
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

// loadDummyFS_Integ builds a real *yongol.Fullstack from a dummy specs dir,
// force-populating ServiceFuncs from ssac.ParseDir (see nestjs counterpart).
func loadDummyFS_Integ(t *testing.T, root string) *yongol.Fullstack {
	t.Helper()
	det, err := yongol.DetectSSOTs(root)
	if err != nil {
		t.Fatalf("DetectSSOTs(%s): %v", root, err)
	}
	fs := yongol.ParseAll(root, det)
	if len(fs.ServiceFuncs) == 0 {
		funcs, _ := pssac.ParseDir(filepath.Join(root, "service"))
		fs.ServiceFuncs = funcs
	}
	return fs
}

// TestGoginSSaCGenerate_Integ runs the gogin SSaC Generate against zenflow
// specs, exercising generateHTTPMethod and generateSubscribeMethod across the
// parsed ServiceFuncs (subscribe vs http branch), plus the converter emit.
func TestGoginSSaCGenerate_Integ(t *testing.T) {
	root := "/home/parkjunwoo/.clari/repos/fullend/dummys/zenflow/try-03/specs"
	if _, err := os.Stat(root); err != nil {
		t.Skipf("dummy specs not present: %v", err)
	}
	fs := loadDummyFS_Integ(t, root)
	if len(fs.ServiceFuncs) == 0 {
		t.Fatal("expected ServiceFuncs from dummy specs")
	}

	out := t.TempDir()
	if err := Generate(fs, out); err != nil {
		t.Fatalf("Generate: %v", err)
	}

	serviceDir := filepath.Join(out, "backend", "internal", "service")
	if _, err := os.Stat(filepath.Join(serviceDir, "server.go")); err != nil {
		t.Errorf("expected server.go: %v", err)
	}
	var goFiles int
	_ = filepath.Walk(serviceDir, func(p string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() && filepath.Ext(p) == ".go" {
			goFiles++
		}
		return nil
	})
	if goFiles < 2 {
		t.Errorf("expected multiple generated .go service files, got %d", goFiles)
	}
}

// TestGoginSSaCGenerate_EmptyEarlyReturn_Integ covers the early return when no
// ServiceFuncs are present.
func TestGoginSSaCGenerate_EmptyEarlyReturn_Integ(t *testing.T) {
	if err := Generate(&yongol.Fullstack{}, t.TempDir()); err != nil {
		t.Fatalf("empty Generate should be a no-op nil, got %v", err)
	}
}

// TestGenerateSubscribeMethod_Integ covers generateSubscribeMethod with a
// hand-built subscribe ServiceFunc. The dummy specs' subscribe handler carries
// a known parse defect (package-qualified @call result) that drops it before it
// reaches Generate, so we construct a minimal one directly here. A single
// @publish sequence keeps the handler free of DB/sqlc dependencies.
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
	if err := generateSubscribeMethod(sf, fs, serviceDir, "github.com/example/app"); err != nil {
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
