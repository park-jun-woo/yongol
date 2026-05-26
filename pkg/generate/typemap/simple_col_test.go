//ff:type feature=gen-typemap type=model
//ff:what simpleCol — 테스트용 ColumnMeta 구현체

package typemap

// simpleCol implements ColumnMeta for testing.
type simpleCol struct {
	raw   string
	enums []string
}
