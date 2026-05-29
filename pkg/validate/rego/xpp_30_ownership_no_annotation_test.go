//ff:func feature=validate type=test control=sequence topic=rego-structural
//ff:what xpp30OwnershipNoAnnotation — resource_owner 참조 + @ownership 누락 검증

package rego

import (
	"strings"
	"testing"

	regoparser "github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXpp30OwnershipNoAnnotation_Unit(t *testing.T) {
	t.Run("no policies returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xpp30OwnershipNoAnnotation(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("no resource_owner usage returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego", Rules: []regoparser.AllowRule{{UsesOwner: false}}},
			},
		}
		diags := xpp30OwnershipNoAnnotation(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("uses owner with ownership annotation passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{
					File:       "auth.rego",
					Rules:      []regoparser.AllowRule{{UsesOwner: true}},
					Ownerships: []regoparser.OwnershipMapping{{Resource: "order"}},
				},
			},
		}
		diags := xpp30OwnershipNoAnnotation(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("uses owner without ownership fires XPP-30", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{
					File:  "auth.rego",
					Rules: []regoparser.AllowRule{{UsesOwner: true}},
				},
			},
		}
		diags := xpp30OwnershipNoAnnotation(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XPP-30]") {
			t.Errorf("expected XPP-30, got %s", diags[0].Message)
		}
	})
}

func TestXpp30OwnershipNoAnnotation(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
