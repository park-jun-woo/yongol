//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectActiveBlocks — Fullstack 상태에 따라 활성 블록만 수집

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// collectActiveBlocks assembles blocks in fixed order and filters to those
// whose Active condition passes (nil Active = always active).
func collectActiveBlocks(fs *yongol.Fullstack, modulePath string) []MainBlock {
	candidates := []MainBlock{
		blockLoggerInit(fs),
		blockEnvHelpers(),
		blockDBInit(fs, modulePath),
		blockJWTSecret(fs),
		blockAuthzInit(fs),
		blockSessionInit(fs),
		blockCacheInit(fs),
		blockFileInit(fs),
		blockServerStruct(fs, modulePath),
		blockQueueInit(fs),
		blockOtelInit(fs, modulePath),
		blockRouter(fs, modulePath),
		blockRequestID(fs, modulePath),
		blockErrorEnvelope(fs, modulePath),
		blockCORS(fs),
		blockPrometheus(fs, modulePath),
		blockSecurityHeaders(fs, modulePath),
		blockCsrf(fs, modulePath),
		blockBodyLimit(fs, modulePath),
		blockRequestValidator(fs, modulePath),
		blockHealth(fs),
		blockRegisterHandlers(fs, modulePath),
		blockAuthInit(fs, modulePath),
		blockGinRun(fs),
	}
	var active []MainBlock
	for _, b := range candidates {
		if b.Active == nil || b.Active(fs) {
			active = append(active, b)
		}
	}
	return active
}
