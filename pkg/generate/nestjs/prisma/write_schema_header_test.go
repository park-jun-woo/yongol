//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestWriteSchemaHeader — Prisma datasource/generator 헤더 블록 출력 검증

package prisma

import (
	"strings"
	"testing"
)

func TestWriteSchemaHeader(t *testing.T) {
	var b strings.Builder
	writeSchemaHeader(&b)
	out := b.String()

	for _, want := range []string{
		"datasource db {",
		`provider = "postgresql"`,
		`url      = env("DATABASE_URL")`,
		"generator client {",
		`provider = "prisma-client-js"`,
	} {
		if !strings.Contains(out, want) {
			t.Errorf("header missing %q\n--- got ---\n%s", want, out)
		}
	}
}
