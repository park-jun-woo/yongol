//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what ParseFile JSON 태그 파싱 테스트 — struct 필드의 json 태그에서 JSONName 추출 검증

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFileJSONTags(t *testing.T) {
	dir := t.TempDir()

	src := `package auth

// @func issueToken
// @description JWT 토큰 발급

type IssueTokenRequest struct {
	Email  string ` + "`" + `json:"email"` + "`" + `
	Role   string ` + "`" + `json:"role"` + "`" + `
	UserID int64  ` + "`" + `json:"user_id"` + "`" + `
}

type IssueTokenResponse struct {
	AccessToken string ` + "`" + `json:"access_token"` + "`" + `
}

func IssueToken(req IssueTokenRequest) (IssueTokenResponse, error) {
	return IssueTokenResponse{AccessToken: "tok"}, nil
}
`
	path := filepath.Join(dir, "issue_token.go")
	os.WriteFile(path, []byte(src), 0644)

	spec, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("ParseFile diagnostics: %v", diags)
	}
	if spec == nil {
		t.Fatal("expected non-nil spec")
	}

	// Request fields should have JSONName.
	if len(spec.RequestFields) != 3 {
		t.Fatalf("RequestFields count = %d, want 3", len(spec.RequestFields))
	}
	for _, f := range spec.RequestFields {
		if f.Name == "UserID" && f.JSONName != "user_id" {
			t.Errorf("UserID.JSONName = %q, want %q", f.JSONName, "user_id")
		}
	}

	// Response field should have JSONName.
	if len(spec.ResponseFields) != 1 {
		t.Fatalf("ResponseFields count = %d, want 1", len(spec.ResponseFields))
	}
	rf := spec.ResponseFields[0]
	if rf.Name != "AccessToken" {
		t.Errorf("Name = %q, want %q", rf.Name, "AccessToken")
	}
	if rf.JSONName != "access_token" {
		t.Errorf("JSONName = %q, want %q", rf.JSONName, "access_token")
	}
}
