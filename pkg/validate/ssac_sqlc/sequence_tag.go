//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what sequenceTag — 시퀀스 Type 을 user-visible SSaC 태그로 변환

package ssac_sqlc

// sequenceTag renders the user-visible SSaC tag for a sequence Type so
// diagnostic messages match the authoring syntax (call / publish / …).
func sequenceTag(t string) string {
	return t
}
