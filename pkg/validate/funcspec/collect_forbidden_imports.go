//ff:func feature=validate type=util control=iteration dimension=1 topic=funcspec-structural
//ff:what collectForbiddenImports — 단일 FuncSpec의 import 중 XFF-41 금지 항목에 대한 Diagnostic 생성

package funcspec

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

func collectForbiddenImports(pkg, name string, line int, imports []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, imp := range imports {
		if !isForbiddenImport(imp) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    pkg + "/" + name + ".go",
			Line:    line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XFF-41] func " + pkg + "." + name + " imports forbidden I/O package " + imp,
			Advice:  "금지된 import 를 제거하고 yongol 가 제공하는 pkg/<X> 를 사용하세요",
		})
	}
	return diags
}
