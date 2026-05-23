//ff:func feature=agent type=command control=sequence
//ff:what generateSplitPathBlock — feature로부터 parameters + requestBody + schema200 분할 생성 후 조립

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func generateSplitPathBlock(feat features.Feature, ddlContent string, cfg Config) (map[string]any, error) {
	params, err := callStepWithRetry(cfg, buildParamsPrompt(feat))
	if err != nil {
		return nil, fmt.Errorf("parameters: %w", err)
	}
	var paramsParsed []any
	if params != "none" && strings.TrimSpace(params) != "" {
		paramsParsed, err = parseParamsYAML(params)
		if err != nil {
			return nil, err
		}
	}

	reqBodyParsed, err := generateReqBody(feat, ddlContent, cfg)
	if err != nil {
		return nil, err
	}

	schema200Raw, err := callStepWithRetry(cfg, buildSchema200Prompt(feat, ddlContent))
	if err != nil {
		return nil, fmt.Errorf("schema200: %w", err)
	}
	schema200Parsed, err := parseSchema200YAML(schema200Raw)
	if err != nil {
		return nil, err
	}

	errorResps := buildErrorResponses(feat)
	block := assemblePathBlock(feat, paramsParsed, reqBodyParsed, schema200Parsed, errorResps)
	return block, nil
}
