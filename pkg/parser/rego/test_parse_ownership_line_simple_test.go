//ff:func feature=policy type=test control=sequence
//ff:what parseOwnershipLine — via 없는 단순 ownership 은 JoinTable/JoinFK 비어있음

package rego

import "testing"

func TestParseOwnershipLine_Simple(t *testing.T) {
	om, ok := parseOwnershipLine("# @ownership gig: gigs.client_id")
	if !ok {
		t.Fatal("expected ok")
	}
	if om.JoinTable != "" || om.JoinFK != "" {
		t.Errorf("join should be empty: %+v", om)
	}
}
