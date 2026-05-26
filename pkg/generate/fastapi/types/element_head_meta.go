//ff:type feature=gen-fastapi type=adapter
//ff:what elementHeadMeta — 배열 element head 를 typemap.ColumnMeta 인터페이스에 맞추는 어댑터

package types

// elementHeadMeta wraps a bare element head token to satisfy
// typemap.ColumnMeta. It reports no check enum and uses the head as the
// raw type, so ClassifyFamily resolves the element's native family.
type elementHeadMeta struct{ head string }
