//ff:func feature=manifest type=parser control=iteration dimension=2
//ff:what parseSentinelInserts — scan raw SQL for top-level INSERT blocks and attach sentinels tagged with `-- @sentinel`
package ddl

import (
	"regexp"
	"strings"
)

// insertIntoRe matches the start of a top-level INSERT, capturing the
// target table name. We accept the common forms used by PostgreSQL
// (optional schema prefix is NOT supported in v1 — keep parser simple).
var insertIntoRe = regexp.MustCompile(`(?i)^\s*INSERT\s+INTO\s+([A-Za-z_][A-Za-z0-9_]*)`)

// onConflictDoNothingRe matches `ON CONFLICT ... DO NOTHING` anywhere in
// the collected INSERT body. Whitespace is flexible; the `...` between
// ON CONFLICT and DO NOTHING (target list, action, etc.) is unrestricted.
var onConflictDoNothingRe = regexp.MustCompile(`(?is)ON\s+CONFLICT\b[^;]*\bDO\s+NOTHING\b`)

// sentinelScanResult carries one INSERT occurrence plus a flag telling
// the caller whether `-- @sentinel` preceded it.
type sentinelScanResult struct {
	Table     string
	SQL       string
	StartLine int
	Annotated bool
}

// SentinelScanResult is the exported mirror of sentinelScanResult for
// downstream packages (pkg/generate/migration) that need to reuse the
// scan logic without duplicating the quote-aware terminator state.
type SentinelScanResult struct {
	Table     string // target table name (raw, un-canonicalised)
	SQL       string // raw SQL including INSERT through final ";"
	StartLine int    // 1-based line number of the INSERT keyword
	Annotated bool   // true when `-- @sentinel` preceded the INSERT
}

// ScanSentinelInserts exposes the parser's sentinel scanner for use by
// other pkg/ modules (e.g. pkg/generate/migration) so the quote-aware
// terminator logic lives in one place.
func ScanSentinelInserts(content string) []SentinelScanResult {
	raw := parseSentinelInserts(content)
	out := make([]SentinelScanResult, 0, len(raw))
	for _, r := range raw {
		out = append(out, SentinelScanResult{
			Table:     r.Table,
			SQL:       r.SQL,
			StartLine: r.StartLine,
			Annotated: r.Annotated,
		})
	}
	return out
}

// SentinelHasOnConflictDoNothing exposes the ON CONFLICT DO NOTHING
// detector so migration / tests can reuse it.
func SentinelHasOnConflictDoNothing(sql string) bool {
	return sentinelHasOnConflictDoNothing(sql)
}

// parseSentinelInserts walks the raw file content line by line, finds
// top-level INSERT statements, and returns each one paired with whether
// the line immediately above it (skipping blank lines) carried the
// `-- @sentinel` annotation. The INSERT body is collected verbatim
// through its terminating `;`. Strings in single quotes are respected so
// a `;` inside a literal does not end the statement prematurely.
func parseSentinelInserts(content string) []sentinelScanResult {
	lines := strings.Split(content, "\n")
	var results []sentinelScanResult
	annotated := false
	i := 0
	for i < len(lines) {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			// blank lines do not reset the annotation (allows one blank between)
			i++
			continue
		}
		if isSentinelAnnotation(trimmed) {
			annotated = true
			i++
			continue
		}
		if m := insertIntoRe.FindStringSubmatch(lines[i]); m != nil {
			table := m[1]
			start := i
			// collect through final `;` at top level (respecting quotes)
			var buf strings.Builder
			j := i
			done := false
			inSingle := false
			for j < len(lines) && !done {
				ln := lines[j]
				if j > i {
					buf.WriteByte('\n')
				}
				buf.WriteString(ln)
				// scan this line for an unquoted `;`
				for k := 0; k < len(ln); k++ {
					ch := ln[k]
					if ch == '\'' {
						// handle doubled `''` as escape inside literal
						if inSingle && k+1 < len(ln) && ln[k+1] == '\'' {
							k++
							continue
						}
						inSingle = !inSingle
						continue
					}
					if ch == ';' && !inSingle {
						// found terminator at position k on line j
						done = true
						// truncate the line after `;` — drop trailing text
						if j == i {
							buf.Reset()
							buf.WriteString(ln[:k+1])
						} else {
							// rebuild buffer with truncated final line
							prev := strings.Join(lines[i:j], "\n")
							buf.Reset()
							buf.WriteString(prev)
							buf.WriteByte('\n')
							buf.WriteString(ln[:k+1])
						}
						break
					}
				}
				j++
			}
			results = append(results, sentinelScanResult{
				Table:     table,
				SQL:       buf.String(),
				StartLine: start + 1,
				Annotated: annotated,
			})
			annotated = false
			i = j
			continue
		}
		// any other non-annotation line invalidates a pending annotation
		annotated = false
		i++
	}
	return results
}

// sentinelHasOnConflictDoNothing reports whether the collected INSERT
// text contains an `ON CONFLICT ... DO NOTHING` clause.
func sentinelHasOnConflictDoNothing(sql string) bool {
	return onConflictDoNothingRe.MatchString(sql)
}
