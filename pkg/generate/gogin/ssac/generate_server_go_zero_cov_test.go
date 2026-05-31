//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestGenerateServerGo_ZeroCov — internal/service/server.go (Server struct) 생성
package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateServerGo_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := generateServerGo(&yongol.Fullstack{}, dir, "example.com/app"); err != nil {
		t.Fatalf("generateServerGo: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "backend", "internal", "service", "server.go"))
	if err != nil {
		t.Fatalf("read server.go: %v", err)
	}
	src := string(data)
	for _, want := range []string{
		"package service",
		"type Server struct {",
		"DB      *pgxpool.Pool",
		"Queries *db.Queries",
		`"example.com/app/internal/db"`,
	} {
		if !strings.Contains(src, want) {
			t.Errorf("missing %q in:\n%s", want, src)
		}
	}
}
