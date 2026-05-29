//ff:func feature=funcspec type=parser control=sequence
//ff:what @func 어노테이션이 파일 중간에 위치할 때도 FuncSpec.Line 이 정확히 채워지는지 검증

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseLine_FuncSpecOffsetAnnotation(t *testing.T) {
	// @func placed deeper to confirm we record the actual annotation line.
	dir := t.TempDir()
	pkgDir := filepath.Join(dir, "billing")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	// Lines:
	// 1: package billing
	// 2: blank
	// 3: import "fmt"
	// 4: blank
	// 5: var _ = fmt.Sprintf
	// 6: blank
	// 7: // @func holdEscrow
	src := `package billing

import "fmt"

var _ = fmt.Sprintf

// @func holdEscrow
// @error 402

type HoldEscrowRequest struct {
	GigID int64
}

type HoldEscrowResponse struct {
	TransactionID int64
}

func HoldEscrow(req HoldEscrowRequest) (HoldEscrowResponse, error) {
	return HoldEscrowResponse{}, nil
}
`
	path := filepath.Join(pkgDir, "hold_escrow.go")
	if err := os.WriteFile(path, []byte(src), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	spec, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("ParseFile diagnostics: %v", diags)
	}
	if spec == nil {
		t.Fatal("expected non-nil spec")
	}

	if spec.Line != 7 {
		t.Errorf("FuncSpec.Line = %d, want 7 (@func annotation line)", spec.Line)
	}
}
