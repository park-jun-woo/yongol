//ff:func feature=policy type=parser control=iteration dimension=1
//ff:what ParsePolicyFile — 단일 .rego 파일에서 Policy 추출 + OPA AST strict parse
package rego

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

var (
	reOwnership = regexp.MustCompile(`^#\s*@ownership\s+(\w+):\s+(\w+)\.(\w+)(?:\s+via\s+(\w+)\.(\w+))?$`)
	reAction    = regexp.MustCompile(`input\.action\s*(?:==\s*"(\w+)"|in\s*\{([^}]+)\})`)
	reResource  = regexp.MustCompile(`input\.resource\s*==\s*"(\w+)"`)
	reOwnerRef  = regexp.MustCompile(`input\.resource_owner`)
	reRoleRef   = regexp.MustCompile(`input\.(?:user|claims)\.role\s*==\s*"(\w+)"`)
	reClaimsRef = regexp.MustCompile(`input\.claims\.(\w+)`)
)

// ParsePolicyFile parses a single .rego file. Two channels run:
//  1. OPA AST strict parse → R-1 ERROR on genuine syntax violations.
//  2. Regex-based metadata extraction (@ownership / allow rule / claims refs)
//     and "malformed @ownership" parse ERRORs.
func ParsePolicyFile(path string) (*Policy, []diagnostic.Diagnostic) {
	f, err := os.Open(path)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "[Rego open] " + err.Error(),
		}}
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "[Rego read] " + err.Error(),
		}}
	}
	content := string(data)

	var diags []diagnostic.Diagnostic
	diags = append(diags, strictParse(path, content)...)

	p := &Policy{File: path}
	lineNo := 0
	for _, line := range strings.Split(content, "\n") {
		lineNo++
		trimmed := strings.TrimSpace(line)
		if om, ok := parseOwnershipLine(trimmed); ok {
			om.SourceLine = lineNo
			p.Ownerships = append(p.Ownerships, om)
			continue
		}
		if looksLikeOwnership(trimmed) {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    lineNo,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("malformed @ownership: %s → 권고: `@ownership <resource>: <table>.<column>` 또는 `@ownership <resource>: <table>.<column> via <join_table>.<join_fk>` 형식", trimmed),
			})
		}
	}

	extractAllowRules(content, p)
	extractClaimsRefs(content, p)
	return p, diags
}
