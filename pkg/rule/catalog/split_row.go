//ff:func feature=rule type=util control=sequence topic=catalog
//ff:what splitRow — markdown 테이블 행을 `|` 기준으로 셀 슬라이스로 분리
package catalog

import "strings"

// splitRow splits a markdown table row on pipe characters, stripping the
// leading and trailing pipes. Rulebook.md doesn't use escaped pipes inside
// cells so a naive split is sufficient.
func splitRow(row string) []string {
	row = strings.TrimSpace(row)
	row = strings.TrimPrefix(row, "|")
	row = strings.TrimSuffix(row, "|")
	return strings.Split(row, "|")
}
