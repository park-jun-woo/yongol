//ff:func feature=agent type=helper control=sequence
//ff:what recordPathOps — YAML path 블록에서 path key 추출 후 pathToOps/opToPath 매핑

package agent

import "gopkg.in/yaml.v3"

func recordPathOps(content, opName string, pathToOps map[string][]string, opToPath map[string]string) {
	var parsed map[string]any
	if err := yaml.Unmarshal([]byte(content), &parsed); err == nil {
		for k := range parsed {
			pathToOps[k] = appendUnique(pathToOps[k], opName)
			opToPath[opName] = k
		}
		return
	}
	wrapped := "paths:\n" + indentText(content, "  ")
	var wrappedParsed struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal([]byte(wrapped), &wrappedParsed); err == nil {
		for k := range wrappedParsed.Paths {
			pathToOps[k] = appendUnique(pathToOps[k], opName)
			opToPath[opName] = k
		}
	}
}
