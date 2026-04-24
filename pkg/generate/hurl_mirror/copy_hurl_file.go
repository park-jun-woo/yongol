//ff:func feature=gen-hurl type=util control=sequence
//ff:what copyHurlFile — src → dst verbatim 복사 (상위 디렉토리 lazy 생성)

package hurl_mirror

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// copyHurlFile copies src to dst byte-for-byte. The parent directory of
// dst is created on demand so nested layouts under specs/tests/ are
// preserved without a pre-walk step.
func copyHurlFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", filepath.Dir(dst), err)
	}
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open %s: %w", src, err)
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create %s: %w", dst, err)
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy %s -> %s: %w", src, dst, err)
	}
	return nil
}
