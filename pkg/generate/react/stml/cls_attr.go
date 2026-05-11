//ff:func feature=stml-gen type=util control=sequence
//ff:what className 문자열을 JSX className 속성으로 변환한다
package stml

import "fmt"

func clsAttr(className string) string {
	if className == "" {
		return ""
	}
	return fmt.Sprintf(` className="%s"`, className)
}
