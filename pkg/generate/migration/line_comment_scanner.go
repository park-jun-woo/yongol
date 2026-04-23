//ff:type feature=migration type=model
//ff:what lineCommentScanner — findLineCommentStart 의 inSQ 플래그 보유 스캐너
package migration

// lineCommentScanner carries the single-quote flag across step calls.
type lineCommentScanner struct {
	inSQ bool
}
