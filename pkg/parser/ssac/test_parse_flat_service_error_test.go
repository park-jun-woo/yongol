//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 플랫 service/ 디렉토리에 .ssac 파일 배치 시 에러 검증

package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseFlatServiceError(t *testing.T) {
	dir := t.TempDir()

	src := `package service

// @get User user = User.FindByEmail({Email: request.Email})
// @response {
//   user: user
// }
func Login() {}
`
	os.WriteFile(filepath.Join(dir, "login.ssac"), []byte(src), 0644)

	_, diags := ParseDir(dir)
	if len(diags) == 0 {
		t.Fatal("expected diagnostic for flat service/ file, got none")
	}
	if !strings.Contains(diags[0].Message, "feature 서브 폴더를 사용하세요") {
		t.Errorf("unexpected diagnostic message: %s", diags[0].Message)
	}
}
