//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestExtractSignatureBasic — param 이름/타입, error 반환 포함한 기본 시그니처 추출

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSignatureBasic(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "svc.go")
	src := "package svc\n\nimport \"context\"\n\nfunc ActivateWorkflow(ctx context.Context, id int64, force bool) (string, error) {\n\treturn \"\", nil\n}\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sig, err := ExtractSignature(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sig.Name != "ActivateWorkflow" {
		t.Errorf("name: got %q want ActivateWorkflow", sig.Name)
	}
	if len(sig.Params) != 3 {
		t.Fatalf("params: got %d (%v) want 3", len(sig.Params), sig.Params)
	}
	wantParams := []FuncParam{
		{Name: "ctx", Type: "context.Context"},
		{Name: "id", Type: "int64"},
		{Name: "force", Type: "bool"},
	}
	for i, p := range wantParams {
		if sig.Params[i] != p {
			t.Errorf("param[%d]: got %+v want %+v", i, sig.Params[i], p)
		}
	}
	if len(sig.Returns) != 2 || sig.Returns[0] != "string" || sig.Returns[1] != "error" {
		t.Errorf("returns: got %v", sig.Returns)
	}
	if !sig.HasErr {
		t.Errorf("expected HasErr = true")
	}
}
