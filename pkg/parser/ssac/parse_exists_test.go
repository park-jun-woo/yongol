//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @exists parse test — verifies Target and Message

package ssac

import "testing"

func TestParseExists(t *testing.T) {
	src := `package service

// @exists existing "already exists"
func CreateCourse(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqExists)
	assertEqual(t, "Target", seq.Target, "existing")
	assertEqual(t, "Message", seq.Message, "already exists")
}
