//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestExtractCORSConfig_ZeroCov(t *testing.T) {
	plan := &ir.BootPlan{ActiveBlocks: []ir.BootBlock{
		{Name: "cors", Active: true, Config: &ir.CORSBootConfig{
			AllowOrigins:     []string{"https://a.com"},
			AllowCredentials: true,
		}},
	}}
	origins, creds := extractCORSConfig(plan)
	if len(origins) != 1 || origins[0] != "https://a.com" || !creds {
		t.Errorf("extractCORSConfig = %v, %v", origins, creds)
	}
	// No cors block → default nil, true
	o2, c2 := extractCORSConfig(&ir.BootPlan{})
	if o2 != nil || !c2 {
		t.Errorf("default = %v, %v", o2, c2)
	}
}
