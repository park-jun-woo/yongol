//ff:func feature=gen-gogin type=util control=sequence
//ff:what shouldKeepImport — 단일 import 라인 유지 여부 결정 (blank / 사용 여부 분기)

package boot

import "strings"

// shouldKeepImport decides whether one import line survives
// filterImportsUsed's sweep. Blank (side-effect) imports follow
// keepBlank; regular imports are kept when their package identifier
// appears in body as `pkg.` — bounded on the left by a non-identifier
// character so `sql.` from `database/sql` is NOT matched by the
// substring inside `otelsql.`. The word-boundary check prevents false
// positives that would otherwise produce unused-import build errors
// when a longer package name contains a shorter one as suffix.
func shouldKeepImport(imp, body string, keepBlank bool) bool {
	if strings.HasPrefix(strings.TrimSpace(imp), "_ ") {
		return keepBlank
	}
	pkg := importIdentifier(imp)
	if pkg == "" {
		return false
	}
	needle := pkg + "."
	idx := 0
	for {
		found := strings.Index(body[idx:], needle)
		if found < 0 {
			return false
		}
		at := idx + found
		if at == 0 || !isIdentByte(body[at-1]) {
			return true
		}
		idx = at + len(needle)
	}
}

// isIdentByte reports whether b could be part of a Go identifier —
// letters, digits, and underscore. Used to enforce the left-side word
// boundary for import usage detection. Only ASCII is considered because
// the generator emits ASCII identifiers only; if the generator ever
// produces unicode identifiers this needs to widen.
func isIdentByte(b byte) bool {
	switch {
	case b >= 'a' && b <= 'z':
		return true
	case b >= 'A' && b <= 'Z':
		return true
	case b >= '0' && b <= '9':
		return true
	case b == '_':
		return true
	}
	return false
}
