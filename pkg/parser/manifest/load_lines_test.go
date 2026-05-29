//ff:func feature=manifest type=parser control=sequence
//ff:what ClaimDef.SourceLine 과 Auth.RolesLines 가 manifest.yaml 의 실제 줄로 채워지는지 검증

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_AuthClaimsAndRolesSourceLine(t *testing.T) {
	dir := t.TempDir()
	// Lines (1-based):
	// 1: apiVersion: yongol/v1
	// 2: kind: Project
	// 3: metadata:
	// 4:   name: testapp
	// 5: backend:
	// 6:   module: github.com/test/testapp
	// 7:   auth:
	// 8:     type: jwt
	// 9:     claims:
	// 10:      UserID: "user_id:int64"
	// 11:      Role: "role"
	// 12:    roles:
	// 13:      - client
	// 14:      - freelancer
	// 15: frontend:
	// 16:   framework: react
	content := "apiVersion: yongol/v1\n" +
		"kind: Project\n" +
		"metadata:\n" +
		"  name: testapp\n" +
		"backend:\n" +
		"  module: github.com/test/testapp\n" +
		"  auth:\n" +
		"    type: jwt\n" +
		"    claims:\n" +
		"      UserID: \"user_id:int64\"\n" +
		"      Role: \"role\"\n" +
		"    roles:\n" +
		"      - client\n" +
		"      - freelancer\n" +
		"frontend:\n" +
		"  framework: react\n"

	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load diagnostics: %v", diags)
	}
	if cfg.Backend.Auth == nil {
		t.Fatal("Backend.Auth is nil")
	}

	uid, ok := cfg.Backend.Auth.Claims["UserID"]
	if !ok {
		t.Fatal("Claims[UserID] missing")
	}
	if uid.SourceLine != 10 {
		t.Errorf("Claims[UserID].SourceLine = %d, want 10", uid.SourceLine)
	}
	role, ok := cfg.Backend.Auth.Claims["Role"]
	if !ok {
		t.Fatal("Claims[Role] missing")
	}
	if role.SourceLine != 11 {
		t.Errorf("Claims[Role].SourceLine = %d, want 11", role.SourceLine)
	}

	if got := cfg.Backend.Auth.RolesLines["client"]; got != 13 {
		t.Errorf("RolesLines[client] = %d, want 13", got)
	}
	if got := cfg.Backend.Auth.RolesLines["freelancer"]; got != 14 {
		t.Errorf("RolesLines[freelancer] = %d, want 14", got)
	}
}
