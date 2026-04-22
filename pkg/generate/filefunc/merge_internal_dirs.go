//ff:func feature=gen-filefunc type=util control=iteration dimension=1
//ff:what mergeInternalDirs — arts/backend/internal 하위 디렉토리 이름을 feature 맵에 추가
package filefunc

import "os"

// mergeInternalDirs scans internalDir for subdirectories and inserts any
// that are missing from dst, assigning a baseline description. Missing
// directory is silently ignored — the generator runs before backend is
// emitted in tests.
func mergeInternalDirs(dst map[string]string, internalDir string) {
	entries, err := os.ReadDir(internalDir)
	if err != nil {
		return
	}
	for _, e := range entries {
		insertDirEntry(dst, e)
	}
}
