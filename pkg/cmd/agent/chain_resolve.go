//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what chainResolve — chain 패키지로 non-SSaC 파일의 operationId 탐색

package agent

import (
	"github.com/park-jun-woo/yongol/pkg/chain"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// chainResolve tries to use the chain package to find an operationId for a non-SSaC file.
func chainResolve(specsDir, relPath string, lookup map[string]features.Feature) (desc, path string) {
	detected, err := yongol.DetectSSOTs(specsDir)
	if err != nil {
		return "", ""
	}
	fs := yongol.ParseAll(specsDir, detected)
	if len(fs.ParseDiagnostics) > 0 {
		return "", ""
	}

	for op, feat := range lookup {
		links, err := chain.Chain(fs, op)
		if err != nil {
			continue
		}
		for _, link := range links {
			if link.File == relPath {
				return feat.Desc, feat.Path
			}
		}
	}
	return "", ""
}
