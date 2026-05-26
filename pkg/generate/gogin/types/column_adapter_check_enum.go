//ff:func feature=gen-gogin type=accessor control=sequence
//ff:what CheckEnum — columnAdapter 의 CHECK 열거 값 반환

package types

func (a columnAdapter) CheckEnum() []string { return a.col.CheckEnum }
