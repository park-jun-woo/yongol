//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanFileForUnmarshalErr — preserved 파일에서 Unmarshal 호출 에러 미처리 수집

package contract

import (
	"go/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanFileForUnmarshalErr parses path and emits a PRV-12 Diagnostic
// for every `json.Unmarshal` / `yaml.Unmarshal` call whose returned
// error is neither checked (err != nil) nor explicitly discarded
// (`_ = ...`). Scoping is per-block: the call's err identifier must
// be guarded before the block ends or errName is reassigned.
func scanFileForUnmarshalErr(path string) []diagnostic.Diagnostic {
	fset, file, err := parseGoFile(path)
	if err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	ast.Inspect(file, func(n ast.Node) bool {
		block, ok := n.(*ast.BlockStmt)
		if !ok {
			return true
		}
		diags = append(diags, unmarshalDiagsInBlock(fset, file, path, block)...)
		return true
	})
	return diags
}
