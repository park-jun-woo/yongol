//ff:func feature=generate type=util control=sequence
//ff:what makeFrontendCopyWalker — filepath.Walk 콜백을 생성하여 src/frontend → arts 경로로 파일 복제
package generate

import (
	"fmt"
	"os"
	"path/filepath"
)

func makeFrontendCopyWalker(srcRoot, dstRoot string) filepath.WalkFunc {
	return func(path string, fi os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if fi.IsDir() {
			if fi.Name() == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		rel, err := filepath.Rel(srcRoot, path)
		if err != nil {
			return err
		}
		if isYongolManaged(rel) {
			return nil
		}
		if !isCopiedExtension(path) {
			return nil
		}
		dst := filepath.Join(dstRoot, rel)
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", filepath.Dir(dst), err)
		}
		return copyUserComponentFile(path, dst)
	}
}
