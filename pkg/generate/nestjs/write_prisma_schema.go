//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writePrismaSchema — DDL → schema.prisma 파일 기록

package nestjs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/prisma"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// writePrismaSchema renders and writes the Prisma schema from DDL tables.
func writePrismaSchema(fs *yongol.Fullstack, backendDir string) error {
	if len(fs.DDLTables) == 0 {
		return nil
	}
	prismaDir := filepath.Join(backendDir, "prisma")
	if err := os.MkdirAll(prismaDir, 0o755); err != nil {
		return fmt.Errorf("mkdir prisma: %w", err)
	}
	schemaContent, err := prisma.RenderSchema(fs.DDLTables)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(prismaDir, "schema.prisma"), []byte(schemaContent), 0o644)
}
