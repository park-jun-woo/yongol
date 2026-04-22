//ff:func feature=generate type=util control=iteration dimension=1
//ff:what copyOPARego — specs/policy/*.rego → arts/backend/policy/*.rego 단순 복사
package generate

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func copyOPARego(fs *yongol.Fullstack, artifactsDir string) error {
	srcDir := filepath.Join(fs.SpecsDir, "policy")
	entries, err := os.ReadDir(srcDir)
	if os.IsNotExist(err) {
		return nil // specs/policy/ 없으면 skip
	}
	if err != nil {
		return fmt.Errorf("read policy dir: %w", err)
	}

	dstDir := filepath.Join(artifactsDir, "backend", "policy")
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return fmt.Errorf("mkdir policy: %w", err)
	}

	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".rego" {
			continue
		}
		src := filepath.Join(srcDir, e.Name())
		dst := filepath.Join(dstDir, e.Name())

		in, err := os.Open(src)
		if err != nil {
			return fmt.Errorf("open %s: %w", src, err)
		}
		out, err := os.Create(dst)
		if err != nil {
			in.Close()
			return fmt.Errorf("create %s: %w", dst, err)
		}
		if _, err := io.Copy(out, in); err != nil {
			in.Close()
			out.Close()
			return fmt.Errorf("copy %s: %w", e.Name(), err)
		}
		in.Close()
		out.Close()
	}
	return nil
}
