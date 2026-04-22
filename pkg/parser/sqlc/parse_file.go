//ff:func feature=orchestrator type=parser control=iteration dimension=1
//ff:what ParseFile — 단일 sqlc 쿼리 파일에서 QuerySpec 목록 추출 (named params 포함)
package sqlc

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// sqlcNameRe matches lines like:  -- name: UserFindByEmail :one
var sqlcNameRe = regexp.MustCompile(`^\s*--\s*name:\s*([A-Za-z_][A-Za-z0-9_]*)(?:\s+:([A-Za-z_]+))?`)

// namedParamRe matches @param_name in SQL body.
var namedParamRe = regexp.MustCompile(`@([a-z_][a-z0-9_]*)`)

// sqlcArgRe matches sqlc.arg(param_name) in SQL body.
var sqlcArgRe = regexp.MustCompile(`sqlc\.arg\(([a-z_][a-z0-9_]*)\)`)

// ParseFile reads one .sql file and returns the QuerySpecs discovered in it.
// Each QuerySpec includes Params extracted from @named_param in the SQL body.
func ParseFile(path string) ([]QuerySpec, []diagnostic.Diagnostic) {
	f, err := os.Open(path)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "cannot open sqlc query file: " + err.Error(),
		}}
	}
	defer f.Close()

	model := modelFromFilename(filepath.Base(path))

	var specs []QuerySpec
	var current *QuerySpec
	var paramSet map[string]bool

	scanner := bufio.NewScanner(f)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		current, paramSet = processSQLCScanLine(line, model, path, lineNo, current, paramSet, &specs)
	}
	// Flush last query
	if current != nil {
		current.Params = sortedKeys(paramSet)
		specs = append(specs, *current)
	}

	if err := scanner.Err(); err != nil {
		return specs, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "cannot read sqlc query file: " + err.Error(),
		}}
	}
	return specs, nil
}
