//ff:func feature=gen-gogin type=util control=sequence topic=import-collect
//ff:what buildEvalImports — @eval emit 시 호출 대상 패키지의 import 경로 수집

package ssac

import (
	"fmt"
	"strings"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func (g *methodGen) buildEvalImports(seq ssacparser.Sequence) []string {
	parts := strings.SplitN(seq.Model, ".", 2)
	pkgName := parts[0]
	if pkgName == "" {
		return nil
	}
	if ssacBuiltinPkgs[pkgName] {
		return []string{fmt.Sprintf(`"github.com/park-jun-woo/ssac/pkg/%s"`, pkgName)}
	}
	return []string{fmt.Sprintf(`"%s/internal/%s"`, g.ModulePath, pkgName)}
}
