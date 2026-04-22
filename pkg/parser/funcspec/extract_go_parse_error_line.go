//ff:func feature=funcspec type=parser control=sequence
//ff:what go/parser 에러에서 첫 에러 위치의 라인 번호를 추출한다
package funcspec

import "go/scanner"

// extractGoParseErrorLine returns the line number of the first parse error
// in a go/parser failure, or 0 if unavailable.
func extractGoParseErrorLine(err error) int {
	if list, ok := err.(scanner.ErrorList); ok && len(list) > 0 {
		return list[0].Pos.Line
	}
	return 0
}
