//ff:func feature=agent type=helper control=sequence
//ff:what parseParamsYAML — LLM 응답을 parameters 배열로 파싱 (wrapping 키 자동 해제)

package agent

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func parseParamsYAML(raw string) ([]any, error) {
	var parsed []any
	if err := yaml.Unmarshal([]byte(raw), &parsed); err == nil {
		return parsed, nil
	}
	var wrapped map[string]any
	if err := yaml.Unmarshal([]byte(raw), &wrapped); err != nil {
		return nil, fmt.Errorf("parse parameters YAML: %w", err)
	}
	inner, ok := wrapped["parameters"]
	if !ok {
		return nil, fmt.Errorf("parse parameters YAML: missing 'parameters' key")
	}
	arr, ok := inner.([]any)
	if !ok {
		return nil, fmt.Errorf("parse parameters YAML: 'parameters' is not an array")
	}
	return arr, nil
}
