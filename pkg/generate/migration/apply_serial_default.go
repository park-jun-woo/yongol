//ff:func feature=migration type=parser control=sequence
//ff:what applySerialDefault — SERIAL 컬럼에 nextval() 기본값 + NOT NULL 부착
package migration

// applySerialDefault attaches a nextval('<t>_<col>_seq') default and
// marks the column NOT NULL when the original DDL used SERIAL and no
// explicit DEFAULT was supplied.
func applySerialDefault(t *Table, col *Column, isSerial bool) {
	if !isSerial || col.Default != "" {
		return
	}
	col.Default = "nextval('" + t.Name + "_" + col.Name + "_seq')"
	col.Nullable = false
}
