//ff:func feature=gen-typemap type=accessor control=sequence
//ff:what CheckEnum — simpleCol 의 CHECK 열거 값 반환 (테스트용)

package typemap

func (c simpleCol) CheckEnum() []string { return c.enums }
