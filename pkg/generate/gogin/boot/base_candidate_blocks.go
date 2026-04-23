//ff:func feature=gen-gogin type=util control=sequence
//ff:what baseCandidateBlocks — main.go 후보 블록 목록 구성 (session/cache/file/queue prepared 분기 포함)

package boot

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// baseCandidateBlocks returns the fixed-order list of MainBlock
// candidates for collectActiveBlocks. Subsystems that have migrated to
// prepared.State (Phase001: session, cache, file, queue) are appended
// only when active; the remainder still reach for raw state via the
// legacy Active predicate.
func baseCandidateBlocks(fs *yongol.Fullstack, p prepared.State, modulePath string) []MainBlock {
	blocks := []MainBlock{
		blockLoggerInit(fs),
		blockEnvHelpers(),
		blockDBInit(fs, modulePath),
		blockJWTSecret(fs),
		blockAuthzInit(fs),
	}
	if p.ActiveBackends.Session != nil {
		blocks = append(blocks, blockSessionInit(*p.ActiveBackends.Session))
	}
	if p.ActiveBackends.Cache != nil {
		blocks = append(blocks, blockCacheInit(*p.ActiveBackends.Cache))
	}
	if p.ActiveBackends.File != nil {
		blocks = append(blocks, blockFileInit(*p.ActiveBackends.File))
	}
	blocks = append(blocks, blockServerStruct(fs, modulePath))
	if p.ActiveBackends.Queue != nil {
		blocks = append(blocks, blockQueueInit(*p.ActiveBackends.Queue, fs.ServiceFuncs))
	}
	blocks = append(blocks,
		blockOtelInit(fs, modulePath),
		blockRouter(fs, modulePath),
		blockRequestID(fs, modulePath),
		blockErrorEnvelope(fs, modulePath),
		blockCORS(fs),
		blockPrometheus(fs, modulePath),
		blockSecurityHeaders(fs, modulePath),
		blockCsrf(p.Auth, modulePath),
		blockBodyLimit(fs, modulePath),
		blockRequestValidator(fs, modulePath),
		blockHealth(fs),
		blockRegisterHandlers(fs, modulePath),
		blockAuthInit(p.Auth, modulePath),
		blockGinRun(fs),
	)
	return blocks
}
