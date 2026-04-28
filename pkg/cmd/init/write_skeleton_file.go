//ff:func feature=cli-init type=util control=sequence
//ff:what writeSkeletonFile — materializes a single skeleton file (rendered or raw)

package cliinit

import (
	"fmt"
	"os"
	"path/filepath"
)

// writeSkeletonFile materializes a single skeleton file. When f.rendered is
// true the contents are run through text/template, otherwise the embedded
// bytes are copied verbatim.
func writeSkeletonFile(targetDir string, data templateData, f skeletonFile) error {
	dest := filepath.Join(targetDir, f.destRel)
	contents, err := loadSkeletonFileContents(data, f)
	if err != nil {
		return err
	}
	if err := os.WriteFile(dest, contents, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", dest, err)
	}
	return nil
}
