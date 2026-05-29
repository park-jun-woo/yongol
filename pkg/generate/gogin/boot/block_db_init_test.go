//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockDBInit — pgxpool 생성 + sqlc Queries 초기화

package boot

import (
	"strings"
	"testing"
)

func TestBlockDBInit(t *testing.T) {
	block := blockDBInit(nil, "example.com/zenflow")
	if block.Name != "db-init" {
		t.Errorf("name = %q, want db-init", block.Name)
	}
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		"ctx, cancelBootstrap := context.WithCancel(context.Background())",
		"pool := initDBPool(ctx)",
		"defer pool.Close()",
		"queries := db.New(pool)",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("blockDBInit lines missing %q, got:\n%s", must, body)
		}
	}
	imp := strings.Join(block.Imports, "\n")
	for _, must := range []string{
		`"github.com/jackc/pgx/v5/pgxpool"`,
		`"example.com/zenflow/internal/db"`,
	} {
		if !strings.Contains(imp, must) {
			t.Errorf("blockDBInit imports missing %q, got:\n%s", must, imp)
		}
	}
	funcs := strings.Join(block.Funcs, "\n")
	if !strings.Contains(funcs, "func initDBPool(ctx context.Context) *pgxpool.Pool {") {
		t.Errorf("must emit initDBPool helper, got:\n%s", funcs)
	}
}
