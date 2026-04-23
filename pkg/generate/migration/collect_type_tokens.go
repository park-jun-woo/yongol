//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what collectTypeTokens — 타입 표현식 토큰 수집 (character varying / timestamp with time zone 등 다단어 지원)
package migration

import "strings"

// collectTypeTokens consumes one or more tokens forming a SQL type
// expression (e.g. "character varying(255)", "timestamp with time
// zone", "varchar(255)", "int[]"). Returns the joined raw type string
// and the remaining tokens.
func collectTypeTokens(toks []string) (string, []string) {
	if len(toks) == 0 {
		return "", nil
	}
	typeParts := []string{toks[0]}
	i := 1
	upper := strings.ToUpper(toks[0])
	if isMultiWordTypeHead(upper) && i < len(toks) {
		i = consumeMultiWordTypeTail(typeParts[:1], toks, i, &typeParts)
	}
	for i < len(toks) && toks[i] == "[]" {
		typeParts[len(typeParts)-1] += "[]"
		i++
	}
	return strings.Join(typeParts, " "), toks[i:]
}
