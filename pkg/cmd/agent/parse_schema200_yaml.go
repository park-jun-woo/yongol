//ff:func feature=agent type=helper control=sequence
//ff:what parseSchema200YAML — LLM 응답을 schema200 맵으로 파싱 (wrapping 키 자동 해제)

package agent

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func parseSchema200YAML(raw string) (map[string]any, error) {
	var parsed map[string]any
	if err := yaml.Unmarshal([]byte(raw), &parsed); err != nil {
		return nil, fmt.Errorf("parse schema200 YAML: %w", err)
	}
	if inner, ok := parsed["schema"]; ok && len(parsed) == 1 {
		if m, ok := inner.(map[string]any); ok {
			return m, nil
		}
	}
	return parsed, nil
}
