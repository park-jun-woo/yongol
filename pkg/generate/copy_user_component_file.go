//ff:func feature=generate type=util control=sequence
//ff:what copyUserComponentFile — src → dst 단일 파일 복제 (io.Copy 기반, defer close)
package generate

import (
	"fmt"
	"io"
	"os"
)

func copyUserComponentFile(src, dst string) error {
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
