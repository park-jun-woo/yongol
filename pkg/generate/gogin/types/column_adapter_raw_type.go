//ff:func feature=gen-gogin type=accessor control=sequence
//ff:what RawType — columnAdapter 의 DDL 원본 타입 반환

package types

func (a columnAdapter) RawType() string { return a.col.RawType }
