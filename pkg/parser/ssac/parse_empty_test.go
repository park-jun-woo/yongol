//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @empty parse test — verifies Target and Message

package ssac

import "testing"

func TestParseEmpty(t *testing.T) {
	src := `package service

// @empty course "course not found"
func GetCourse(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqEmpty)
	assertEqual(t, "Target", seq.Target, "course")
	assertEqual(t, "Message", seq.Message, "course not found")
}
