//ff:func feature=gen-react type=test control=sequence
//ff:what writeOpenapiTsStub fallback api.d.ts 생성·사유 주석 포함 검증
package react

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteOpenapiTsStub(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "api.d.ts")
	reason := errors.New("openapi-typescript not installed")

	writeOpenapiTsStub(dest, reason)

	data, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "openapi-typescript could not run: openapi-typescript not installed")
	assertContains(t, content, "export type paths = Record<string, any>")
	assertContains(t, content, "export type operations = Record<string, any>")
}
