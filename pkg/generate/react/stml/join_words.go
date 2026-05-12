//ff:func feature=stml-gen type=util control=sequence
//ff:what 단어 슬라이스를 공백으로 연결한다
package stml

import "strings"

// joinWords joins a slice of words with spaces.
func joinWords(words []string) string {
	return strings.Join(words, " ")
}
