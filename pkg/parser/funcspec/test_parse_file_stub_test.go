//ff:func feature=funcspec type=parser control=sequence
//ff:what ParseFile stub 함수 감지 테스트 — 빈 반환문만 있으면 HasBody=false

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFileStub(t *testing.T) {
	dir := t.TempDir()

	src := `package billing

// @func calculateRefund
// @description 환불 금액을 계산한다

type CalculateRefundRequest struct {
	Amount int
}

type CalculateRefundResponse struct {
	Refund int
}

func CalculateRefund(req CalculateRefundRequest) (CalculateRefundResponse, error) {
	return CalculateRefundResponse{}, nil
}
`
	path := filepath.Join(dir, "calculate_refund.go")
	os.WriteFile(path, []byte(src), 0644)

	spec, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("ParseFile diagnostics: %v", diags)
	}
	if spec == nil {
		t.Fatal("expected non-nil spec")
	}

	if spec.HasBody {
		t.Error("HasBody = true, want false (stub)")
	}
}
