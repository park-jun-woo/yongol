//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what 디렉토리 내 모든 .go 파일에서 구조체 타입과 필드를 수집하고, Go parse 실패는 Diagnostic 으로 보고한다
package funcspec

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// collectPackageTypes parses all .go files in dir (non-recursive)
// and returns a map of struct name to fields.
//
// Diagnostic 반환 규약:
//   - dir 이 존재하지 않으면 SILENT-OK (companion 타입 디렉토리 부재는 정상).
//   - 그 외 ReadDir 에러는 Diagnostic 1 건 + 빈 result.
//   - 개별 파일의 parser.ParseFile 실패는 Diagnostic 1 건으로 기록하고
//     나머지 파일은 계속 수집한다 (partial success).
func collectPackageTypes(dir string) (map[string][]Field, []diagnostic.Diagnostic) {
	result := make(map[string][]Field)
	var diags []diagnostic.Diagnostic

	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return result, nil
		}
		return result, []diagnostic.Diagnostic{{
			File:    dir,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "cannot read funcspec type dir: " + err.Error(),
		}}
	}

	fset := token.NewFileSet()
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    extractGoParseErrorLine(err),
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "Go parse failed: " + err.Error(),
			})
			continue
		}
		collectStructsFromFile(f, result)
	}
	return result, diags
}
