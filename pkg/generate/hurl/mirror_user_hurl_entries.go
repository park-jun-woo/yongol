//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what mirrorUserHurlEntries — specsDir/tests/ 항목 중 user hurl 만 dstDir 로 복사

package hurl

import (
	"os"
	"path/filepath"
)

// mirrorUserHurlEntries iterates directory entries and copies every
// user-authored hurl file into dstDir. dstDir is created lazily on the
// first match so projects without user hurls do not leave an empty
// tests/ directory behind. Extracted from mirrorUserHurlFiles to keep
// the main func at control=sequence.
func mirrorUserHurlEntries(srcDir, dstDir string, entries []os.DirEntry) error {
	var mirrored bool
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !isUserHurlName(name) {
			continue
		}
		var err error
		mirrored, err = ensureMirrorDstDir(dstDir, mirrored)
		if err != nil {
			return err
		}
		src := filepath.Join(srcDir, name)
		dst := filepath.Join(dstDir, name)
		if err := copyHurlFile(src, dst); err != nil {
			return err
		}
	}
	return nil
}
