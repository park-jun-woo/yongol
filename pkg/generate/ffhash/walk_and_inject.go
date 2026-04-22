//ff:func feature=gen-ffhash type=util control=sequence
//ff:what WalkAndInject — 디렉토리를 재귀 순회하며 .go 파일에 //ff:checked 삽입 (제외 경로는 건너뜀)

package ffhash

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// WalkAndInject walks root recursively, reads every ".go" file whose
// relative path does not match an entry in skipRelPrefixes, and
// rewrites it via InjectCheckedLine when the annotation block is
// present.
//
// The function is a safety net for write sites that do not yet route
// through WriteManyFiles or writeSplitUnit — running it once at the
// end of `gogin.Generate` guarantees every yongol-owned artifact
// carries a `//ff:checked llm=yongol-gen hash=<sha>` line, independent
// of which emit helper produced the file.
//
// skipRelPrefixes entries are matched against the path relative to
// root, using filepath.ToSlash so callers may use forward slashes on
// every OS. Both directories and individual files may be listed.
//
// Errors from os.ReadFile / os.WriteFile abort the walk immediately —
// emitting some files with hash and others without would silently
// leave A7 drift in the artifact, which is exactly what Phase007
// exists to prevent.
func WalkAndInject(root string, skipRelPrefixes []string) error {
	info, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if !info.IsDir() {
		return nil
	}
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		rel, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return relErr
		}
		slashRel := filepath.ToSlash(rel)
		for _, skip := range skipRelPrefixes {
			if slashRel == skip || strings.HasPrefix(slashRel, strings.TrimSuffix(skip, "/")+"/") {
				return nil
			}
		}
		return rewriteOne(path)
	})
}
