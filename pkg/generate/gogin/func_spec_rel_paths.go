//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what funcSpecRelPaths — specs/func/<pkg>/ 디렉토리를 읽어 "internal/<pkg>" 상대경로 리스트 반환

package gogin

import (
	"os"
	"path/filepath"
)

// funcSpecRelPaths returns the list of "internal/<pkg>" relative paths
// that must be skipped during the ff-checked injection walk. Entries
// come from specs/func/<pkg>/ — the directories copyFuncSpecs mirrors
// into backend/internal/<pkg>/. A non-existent specs/func directory
// returns an empty slice (not an error) so projects without func specs
// continue to work.
//
// Only immediate sub-directories are listed; the walker uses prefix
// matching to exclude every file below each entry, so there is no need
// to recurse here.
func funcSpecRelPaths(specsDir string) ([]string, error) {
	funcDir := filepath.Join(specsDir, "func")
	entries, err := os.ReadDir(funcDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var skip []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		skip = append(skip, "internal/"+e.Name())
	}
	return skip, nil
}
