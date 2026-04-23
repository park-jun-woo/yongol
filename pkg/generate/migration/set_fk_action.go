//ff:func feature=migration type=util control=selection
//ff:what setFKAction — action=DELETE/UPDATE 에 따라 fk.OnDelete/OnUpdate 설정
package migration

// setFKAction assigns val to fk.OnDelete or fk.OnUpdate depending on
// whether action is "DELETE" or "UPDATE".
func setFKAction(fk *ForeignKey, action, val string) {
	switch action {
	case "DELETE":
		fk.OnDelete = val
	case "UPDATE":
		fk.OnUpdate = val
	}
}
