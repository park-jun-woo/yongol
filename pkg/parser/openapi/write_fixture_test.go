//ff:func feature=manifest type=parser control=sequence
//ff:what writeFixture — line_index 테스트용 openapi.yaml 픽스처를 임시 디렉터리에 기록해 경로 반환

package openapi

import (
	"os"
	"path/filepath"
	"testing"
)

// fixture content with intentionally specific line numbers so the test can
// assert exact mappings.
const lineIndexFixture = `openapi: 3.0.3
info:
  title: t
  version: "0"
components:
  schemas:
    User:
      type: object
      properties:
        id: { type: integer, format: int64 }
        email: { type: string }
        name: { type: string }
paths:
  /login:
    post:
      operationId: Login
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [email, password]
              properties:
                email: { type: string, format: email, maxLength: 255 }
                password: { type: string, minLength: 8 }
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
                type: object
                properties:
                  access_token: { type: string }
  /users/me:
    get:
      operationId: GetCurrentUser
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
                type: object
                properties:
                  user: { type: object }
`

func writeFixture(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "openapi.yaml")
	if err := os.WriteFile(path, []byte(lineIndexFixture), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}
