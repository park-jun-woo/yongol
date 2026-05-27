//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what extractCORSConfig — BootPlan CORS 블록에서 origin/credentials 추출

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// extractCORSConfig pulls allowed origins and credentials from the CORS
// boot block config.
func extractCORSConfig(plan *ir.BootPlan) (origins []string, credentials bool) {
	for _, block := range plan.ActiveBlocks {
		if block.Name != "cors" || !block.Active {
			continue
		}
		if cfg, ok := block.Config.(*ir.CORSBootConfig); ok && cfg != nil {
			return cfg.AllowOrigins, cfg.AllowCredentials
		}
	}
	return nil, true
}
