//ff:func feature=funcspec type=parser control=iteration dimension=2
//ff:what CollectAllPackageTypes — rootDir 하위 모든 패키지 디렉토리에서 struct 타입+필드를 재귀 수집

package funcspec

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// CollectAllPackageTypes recursively walks rootDir and collects struct
// types from every subdirectory containing .go files. Each directory is
// processed by collectPackageTypes (non-recursive).
//
// Returns map[pkgName]map[typeName][]Field where pkgName is
// filepath.Base(dir) — matching FuncSpec.Package. Diagnostics from
// individual directories are accumulated and returned alongside.
func CollectAllPackageTypes(rootDir string) (map[string]map[string][]Field, []diagnostic.Diagnostic) {
	result := make(map[string]map[string][]Field)
	var diags []diagnostic.Diagnostic

	// Track directories that contain .go files.
	goDirs := make(map[string]bool)

	filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    0,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "walk error: " + err.Error(),
			})
			return nil
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") && !strings.HasSuffix(d.Name(), "_test.go") {
			goDirs[filepath.Dir(path)] = true
		}
		return nil
	})

	for dir := range goDirs {
		typeMap, dd := collectPackageTypes(dir)
		diags = append(diags, dd...)
		if len(typeMap) == 0 {
			continue
		}
		pkgName := filepath.Base(dir)
		if existing, ok := result[pkgName]; ok {
			// Merge into existing map (multiple dirs with same base name
			// is unlikely but handled defensively).
			for k, v := range typeMap {
				existing[k] = v
			}
		} else {
			result[pkgName] = typeMap
		}
	}
	return result, diags
}
