//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldSSaC — features.yaml ops로부터 SSaC 서비스 파일 자동 생성

package agent

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/docs"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// scaffoldSSaC generates specs/service/{domain}/{Op}.ssac from features.yaml.
// Each feature op produces one LLM call that returns one SSaC file.
// Existing files are skipped to protect user modifications.
func scaffoldSSaC(specsDir string, ff *features.FeaturesFile, openapiContent string, llmFn LLMCallFunc, cfg Config, out io.Writer) (int, error) {
	if len(ff.Features) == 0 {
		return 0, nil
	}

	systemDoc, err := docs.FS.ReadFile("ssac.md")
	if err != nil {
		return 0, fmt.Errorf("read ssac.md docs: %w", err)
	}
	systemPrompt := string(systemDoc)

	count := 0
	for _, feat := range ff.Features {
		created, err := scaffoldSSaCFeature(specsDir, feat, openapiContent, systemPrompt, cfg, out)
		if err != nil {
			return 0, err
		}
		if created {
			count++
		}
	}

	return count, nil
}
