//ff:func feature=migration type=parser control=sequence
//ff:what applyAlterFKOnActions — ALTER TABLE ADD FK tail 에서 ON DELETE/UPDATE 액션 파싱
package migration

import "strings"

// applyAlterFKOnActions sets fk.OnDelete / fk.OnUpdate from the tail
// portion of an ALTER TABLE ADD FOREIGN KEY statement.
func applyAlterFKOnActions(fk *ForeignKey, tail string) {
	u := strings.ToUpper(tail)
	fk.OnDelete = matchFKAction(u, "ON DELETE")
	fk.OnUpdate = matchFKAction(u, "ON UPDATE")
}
