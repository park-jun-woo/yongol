//ff:func feature=rule type=util control=iteration dimension=2
//ff:what splitPascal — PascalCase 문자열을 토큰으로 분해 ("OrgIdBox" → ["Org","Id","Box"])
package rule

// splitPascal splits PascalCase into tokens. "OrgIdBox" → ["Org", "Id", "Box"].
// Runs of uppercase stay together ("URL" → ["URL"]).
func splitPascal(s string) []string {
	var out []string
	i := 0
	for i < len(s) {
		if !isUpper(s[i]) {
			// leading lowercase (shouldn't happen for PascalCase; handle gracefully)
			j := i + 1
			for j < len(s) && !isUpper(s[j]) {
				j++
			}
			out = append(out, s[i:j])
			i = j
			continue
		}
		// uppercase run
		j := i + 1
		for j < len(s) && isUpper(s[j]) {
			j++
		}
		if j == len(s) {
			out = append(out, s[i:j])
			i = j
			continue
		}
		// at j-1 is last upper, at j is lower. The last upper belongs to next word.
		if j-i > 1 {
			out = append(out, s[i:j-1])
			i = j - 1
		}
		// now i is uppercase start; take the rest of lower run as the word
		j = i + 1
		for j < len(s) && !isUpper(s[j]) {
			j++
		}
		out = append(out, s[i:j])
		i = j
	}
	return out
}
