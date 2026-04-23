//ff:func feature=policy type=test control=sequence
//ff:what parseOwnershipLine — via 절이 포함된 2단계 ownership 문자열 파싱

package rego

import "testing"

func TestParseOwnershipLine_WithVia(t *testing.T) {
	om, ok := parseOwnershipLine("# @ownership proposal: proposals.freelancer_id via gigs.client_id")
	if !ok {
		t.Fatal("expected ok")
	}
	if om.Resource != "proposal" || om.Table != "proposals" || om.Column != "freelancer_id" {
		t.Errorf("base = %+v", om)
	}
	if om.JoinTable != "gigs" || om.JoinFK != "client_id" {
		t.Errorf("join = %+v", om)
	}
}
