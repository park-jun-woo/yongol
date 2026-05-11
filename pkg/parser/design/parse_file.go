//ff:func feature=frontend type=parser control=sequence
//ff:what DESIGN.md 파일을 읽어 YAML front matter + body 헤딩을 파싱하여 DesignSpec 반환
package design

import (
	"os"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

var reYAMLLine = regexp.MustCompile(`line (\d+)`)

// ParseFile reads and parses a DESIGN.md file at the given path.
// Returns the DesignSpec and any parse-phase diagnostics.
func ParseFile(path string) (*DesignSpec, []diagnostic.Diagnostic) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "DESIGN.md read error: " + err.Error(),
		}}
	}

	yamlPart, body, err := parseFrontMatter(data)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: err.Error(),
		}}
	}

	var spec DesignSpec
	if err := yaml.Unmarshal(yamlPart, &spec); err != nil {
		line := 0
		if m := reYAMLLine.FindStringSubmatch(err.Error()); len(m) == 2 {
			line, _ = strconv.Atoi(m[1])
		}
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "DESIGN.md YAML parse error: " + err.Error(),
		}}
	}

	spec.File = path
	spec.Headings = parseHeadings(body)

	return &spec, nil
}
