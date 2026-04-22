//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what generateCurrentUser — internal/model/current_user.go 생성 (JSON 태그 포함)

package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateCurrentUser writes internal/model/current_user.go with a CurrentUser
// struct whose fields match the manifest claims. JSON tags use the manifest
// claim keys (snake_case) so that ssac/pkg/authz can pass the struct directly
// to OPA as `input.claims` without a separate Claim type or map conversion.
func generateCurrentUser(artifactsDir string, fields []ClaimField) error {
	dir := filepath.Join(artifactsDir, "backend", "internal", "model")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	var sb strings.Builder
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Type: ffannot.TypeAnnot{Feature: "model", Type: "model"},
		What: "CurrentUser — BearerAuth가 추출한 JWT 클레임 struct (authz.Check input.claims)",
	}))
	sb.WriteString("package model\n\n")
	sb.WriteString("// CurrentUser holds JWT claims extracted by BearerAuth middleware.\n")
	sb.WriteString("// JSON tags mirror manifest claim keys so authz.Check can pass it\n")
	sb.WriteString("// straight to OPA as input.claims via json.Marshal.\n")
	sb.WriteString("type CurrentUser struct {\n")
	for _, f := range fields {
		sb.WriteString(fmt.Sprintf("\t%s %s `json:%q`\n", f.Name, f.GoType, f.Key))
	}
	sb.WriteString("}\n")

	return os.WriteFile(filepath.Join(dir, "current_user.go"), []byte(sb.String()), 0o644)
}
