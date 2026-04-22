//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @auth parse test — verifies Action, Resource, Message, and Inputs

package ssac

import "testing"

func TestParseAuth(t *testing.T) {
	src := `package service

// @auth "delete" "project" {id: project.ID, owner: project.OwnerID} "access denied"
func DeleteProject(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqAuth)
	assertEqual(t, "Action", seq.Action, "delete")
	assertEqual(t, "Resource", seq.Resource, "project")
	assertEqual(t, "Message", seq.Message, "access denied")
	assertEqual(t, "Inputs[id]", seq.Inputs["id"], "project.ID")
	assertEqual(t, "Inputs[owner]", seq.Inputs["owner"], "project.OwnerID")
}
