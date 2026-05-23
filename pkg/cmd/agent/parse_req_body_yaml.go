//ff:func feature=agent type=helper control=sequence
//ff:what parseReqBodyYAML — LLM 응답을 requestBody 맵으로 파싱 (wrapping 키 자동 해제)

package agent

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func parseReqBodyYAML(raw string) (map[string]any, error) {
	var parsed map[string]any
	if err := yaml.Unmarshal([]byte(raw), &parsed); err != nil {
		return nil, fmt.Errorf("parse requestBody YAML: %w", err)
	}
	if inner, ok := parsed["requestBody"]; ok && len(parsed) == 1 {
		if m, ok := inner.(map[string]any); ok {
			return m, nil
		}
	}
	return parsed, nil
}
