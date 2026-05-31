//ff:func feature=gen-gogin type=test-helper control=sequence
//ff:what authClaimsState — bearer 모드 단일 ID 클레임 prepared.State 생성 헬퍼
package auth

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// authClaimsState returns a prepared.State with bearer auth and a single typed
// ID claim.
func authClaimsState() prepared.State {
	return prepared.State{Auth: prepared.Auth{
		Present: true,
		Mode:    "bearer",
		Raw: &manifest.Auth{
			Claims: map[string]manifest.ClaimDef{
				"ID": {Key: "user_id", GoType: "int64", Typed: true},
			},
		},
	}}
}
