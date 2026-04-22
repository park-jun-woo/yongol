//ff:func feature=gen-filefunc type=util control=sequence
//ff:what insertDirEntry — DirEntry 가 디렉토리이고 기존 키에 없을 때만 feature 맵에 삽입
package filefunc

import "os"

// insertDirEntry adds the given directory entry to dst when e is a directory
// and dst does not already contain the key. Files are skipped.
func insertDirEntry(dst map[string]string, e os.DirEntry) {
	if !e.IsDir() {
		return
	}
	name := e.Name()
	if _, ok := dst[name]; ok {
		return
	}
	dst[name] = resolveFeatureDescription(name, "")
}
