//ff:func feature=gen-hurl type=util control=sequence
//ff:what copyHurlFile — src → dst verbatim 파일 복사

package hurl

import (
	"fmt"
	"io"
	"os"
)

// copyHurlFile copies src → dst verbatim. Truncates dst if it exists.
func copyHurlFile(src, dst string) error {
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
