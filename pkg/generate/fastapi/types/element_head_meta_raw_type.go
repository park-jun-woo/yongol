//ff:func feature=gen-fastapi type=adapter control=sequence
//ff:what RawType — elementHeadMeta 의 head 토큰 반환

package types

// RawType returns the element head token as the raw type.
func (m elementHeadMeta) RawType() string { return m.head }
