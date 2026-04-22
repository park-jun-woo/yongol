//ff:func feature=orchestrator type=parser control=iteration dimension=1
//ff:what 디렉토리 내 모든 .rego 파일을 opa/ast로 파싱
package rego

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/open-policy-agent/opa/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseDir parses all .rego files in the given directory using opa/ast.
func ParseDir(dir string) ([]*ast.Module, []diagnostic.Diagnostic) {
	var modules []*ast.Module
	var diags []diagnostic.Diagnostic

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    dir,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "cannot read policy directory: " + err.Error(),
		}}
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".rego") {
			continue
		}
		filePath := filepath.Join(dir, e.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:    filePath,
				Line:    0,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "cannot read rego file: " + err.Error(),
			})
			continue
		}
		module, err := ast.ParseModule(e.Name(), string(data))
		if err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:    filePath,
				Line:    extractErrorLine(err),
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "rego parse error: " + err.Error(),
			})
			continue
		}
		modules = append(modules, module)
	}
	return modules, diags
}
