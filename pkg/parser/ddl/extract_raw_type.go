//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what extractRawType — parts[1:] 로부터 RawType 토큰 (다중 단어 PG 타입 포함) 과 소비 토큰 수 산출

package ddl

import "strings"

// extractRawType inspects the leading tokens of a column definition's
// type region and returns the verbatim multi-word type when matched,
// otherwise the single token. The second value is the number of input
// tokens consumed (always ≥ 1 unless tokens is empty).
//
// A trailing parameter list "(N)" or "(N,M)" attached to the LAST
// matched token is preserved on that token (e.g. "CHARACTER VARYING(255)"
// is recognised as one match consuming 2 tokens, with the joined RawType
// "CHARACTER VARYING(255)"). A trailing comma on the last token is left
// in place — parse_column_def strips it after this helper returns.
func extractRawType(tokens []string) (string, int) {
	if len(tokens) == 0 {
		return "", 0
	}
	upper := make([]string, len(tokens))
	for i, t := range tokens {
		upper[i] = strings.ToUpper(t)
	}
	for _, head := range multiTokenHeads {
		if n := matchMultiTokenHead(upper, head); n > 0 {
			joined := strings.Join(tokens[:n], " ")
			return strings.ToUpper(joined), n
		}
	}
	return upper[0], 1
}
