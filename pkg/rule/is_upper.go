//ff:func feature=rule type=util control=sequence
//ff:what isUpper — ASCII 대문자 여부 검사 (A~Z)
package rule

func isUpper(b byte) bool { return b >= 'A' && b <= 'Z' }
