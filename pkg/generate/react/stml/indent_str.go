//ff:func feature=stml-gen type=util control=sequence
//ff:what 지정된 수만큼 공백 문자열을 생성한다
package stml

import "strings"

func indentStr(n int) string {
	return strings.Repeat(" ", n)
}
