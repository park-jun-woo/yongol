//ff:func feature=agent type=helper control=sequence
//ff:what generateReqBody — requestBody LLM 호출 + 파싱

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func generateReqBody(feat features.Feature, ddlContent string, cfg Config) (map[string]any, error) {
	method := httpMethodFromOp(feat.Op)
	if !needsRequestBody(method) {
		return nil, nil
	}

	reqBodyRaw, err := callStepWithRetry(cfg, buildRequestBodyPrompt(feat, ddlContent))
	if err != nil {
		return nil, fmt.Errorf("requestBody: %w", err)
	}
	if reqBodyRaw == "none" || strings.TrimSpace(reqBodyRaw) == "" {
		return nil, nil
	}
	return parseReqBodyYAML(reqBodyRaw)
}
