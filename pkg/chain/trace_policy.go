//ff:func feature=chain type=util control=iteration dimension=3
//ff:what tracePolicy finds Rego policies referenced by @auth sequences.
package chain

import (
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func tracePolicy(sf *ssac.ServiceFunc, policies []rego.Policy, specsDir string) []Link {
	resources := map[string]bool{}
	actions := map[string]bool{}
	for _, seq := range sf.Sequences {
		if seq.Type != "auth" {
			continue
		}
		if seq.Resource != "" {
			resources[seq.Resource] = true
		}
		if seq.Action != "" {
			actions[seq.Action] = true
		}
	}

	if len(resources) == 0 {
		slog.Debug("chain.tracePolicy: no @auth resources in SSaC function", "operationId", sf.Name)
		return nil
	}

	var links []Link
	seen := map[string]bool{}
	for _, p := range policies {
		for _, rule := range p.Rules {
			if !resources[rule.Resource] {
				continue
			}
			relPath, _ := filepath.Rel(specsDir, p.File)
			if relPath == "" {
				relPath = p.File
			}
			if seen[relPath] {
				continue
			}
			seen[relPath] = true

			line := grepLine(p.File, rule.Resource)
			var actList []string
			for _, a := range rule.Actions {
				if actions[a] {
					actList = append(actList, a)
				}
			}
			summary := "resource: " + rule.Resource
			if len(actList) > 0 {
				summary += " [" + strings.Join(actList, ", ") + "]"
			}
			links = append(links, Link{
				Kind:    "Rego",
				File:    relPath,
				Line:    line,
				Summary: summary,
			})
		}
	}
	return links
}
