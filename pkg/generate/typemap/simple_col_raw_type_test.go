//ff:func feature=gen-typemap type=accessor control=sequence
//ff:what RawType — simpleCol 의 DDL 원본 타입 반환 (테스트용)

package typemap

func (c simpleCol) RawType() string { return c.raw }
