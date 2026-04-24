//ff:func feature=validate type=rule control=iteration dimension=2 topic=ddl-structural
//ff:what D-9 — top-level INSERT without `-- @sentinel` annotation is forbidden

package ddl

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// insertIntoLineRe captures the target table name from an INSERT INTO line
// (used only for error messages). Duplicated locally to keep validate self-
// contained (parser owns the canonical scanner).
var insertIntoLineRe = regexp.MustCompile(`(?i)^\s*INSERT\s+INTO\s+([A-Za-z_][A-Za-z0-9_]*)`)

// d09TopLevelInsertWithoutSentinel validates D-9: any top-level INSERT
// inside a DDL file (specs/db/*.sql) must be preceded by a `-- @sentinel`
// annotation, otherwise the statement would be silently dropped by the
// migration emitter.
func d09TopLevelInsertWithoutSentinel(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, f := range files {
		for _, r := range scanInsertsWithAnnotations(f.content) {
			if r.Annotated {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:  f.path,
				Line:  r.StartLine,
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelError,
				Message: fmt.Sprintf(
					"[D-9] Top-level INSERT into %q has no `-- @sentinel` annotation.",
					r.Table),
				Advice: "Add `-- @sentinel` directly above the INSERT, or move the INSERT out of db/*.sql.",
			})
		}
	}
	return diags
}

// scanInsertsWithAnnotations is a thin adapter that calls the parser's
// sentinel scanner. We depend on the regex-driven scanner living in the
// parser package (pkg/parser/ddl) through a private helper to avoid
// duplicating the quote-aware terminator logic.
//
// To keep this package free of a cycle on parser/ddl we re-implement the
// light-weight scan here. The parser has its own authoritative copy; both
// must agree on what "top-level INSERT" means.
func scanInsertsWithAnnotations(content string) []insertScan {
	lines := strings.Split(content, "\n")
	var results []insertScan
	annotated := false
	i := 0
	for i < len(lines) {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			i++
			continue
		}
		if trimmed == "-- @sentinel" || strings.TrimSpace(strings.TrimPrefix(trimmed, "--")) == "@sentinel" {
			annotated = true
			i++
			continue
		}
		if m := insertIntoLineRe.FindStringSubmatch(lines[i]); m != nil {
			table := m[1]
			start := i
			// collect INSERT body through unquoted `;`
			var buf strings.Builder
			j := i
			done := false
			inSingle := false
			for j < len(lines) && !done {
				ln := lines[j]
				if j > i {
					buf.WriteByte('\n')
				}
				for k := 0; k < len(ln); k++ {
					ch := ln[k]
					if ch == '\'' {
						if inSingle && k+1 < len(ln) && ln[k+1] == '\'' {
							k++
							continue
						}
						inSingle = !inSingle
						continue
					}
					if ch == ';' && !inSingle {
						done = true
						break
					}
				}
				buf.WriteString(ln)
				j++
			}
			results = append(results, insertScan{
				Table:     table,
				SQL:       buf.String(),
				StartLine: start + 1,
				Annotated: annotated,
			})
			annotated = false
			i = j
			continue
		}
		annotated = false
		i++
	}
	return results
}

type insertScan struct {
	Table     string
	SQL       string
	StartLine int
	Annotated bool
}
