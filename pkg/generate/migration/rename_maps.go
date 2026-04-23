//ff:func feature=migration type=util control=iteration dimension=1
//ff:what renameMaps — Hints.RenameTables 로부터 prev→curr/curr→prev 매핑 맵 2개 생성
package migration

// renameMaps returns two lookup maps derived from hints.RenameTables:
// renamed[prev_name] = new_name, renamedRev[new_name] = prev_name.
func renameMaps(hints *Hints) (map[string]string, map[string]string) {
	renamed := map[string]string{}
	renamedRev := map[string]string{}
	if hints == nil {
		return renamed, renamedRev
	}
	for _, r := range hints.RenameTables {
		renamed[r.From] = r.To
		renamedRev[r.To] = r.From
	}
	return renamed, renamedRev
}
