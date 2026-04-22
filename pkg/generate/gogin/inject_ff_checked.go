//ff:func feature=gen-gogin type=util control=sequence
//ff:what injectFFChecked — backend/ 재귀 순회해 yongol-소유 Go 파일에 //ff:checked 주입 (copyFuncSpecs 대상 + sqlc 보존 파일 제외)

package gogin

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ffhash"
)

// injectFFChecked walks backendDir recursively and injects (or refreshes)
// the `//ff:checked llm=yongol-gen hash=<sha>` annotation on every
// yongol-owned Go file. It is invoked at the tail of `gogin.Generate`
// as a safety net so write sites that do not route through
// WriteManyFiles or the splitter still end up with a hash line.
//
// The following paths are excluded:
//
//   - internal/<pkg>/ subtrees mirrored from specs/func/ via copyFuncSpecs
//     (user-authored func specs — yongol does not own or annotate them)
//   - internal/db/querier.go, internal/db/db.go (sqlc core files preserved
//     by splitter.isPreservedFile — leaving them untouched matches the
//     splitter contract)
//
// A missing specsDir/func directory yields an empty skip list, which is
// the correct behaviour for projects without func specs.
func injectFFChecked(backendDir, specsDir string) error {
	skip, err := funcSpecRelPaths(specsDir)
	if err != nil {
		return err
	}
	skip = append(skip, "internal/db/querier.go", "internal/db/db.go")
	return ffhash.WalkAndInject(backendDir, skip)
}
