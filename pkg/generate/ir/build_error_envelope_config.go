//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildErrorEnvelopeConfig -- manifest.backend.error → ErrorEnvelopeConfig 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildErrorEnvelopeConfig always returns a config (error envelope is
// always active).
func buildErrorEnvelopeConfig(fs *yongol.Fullstack) *ErrorEnvelopeConfig {
	return &ErrorEnvelopeConfig{
		ExposeInternalError: resolveExposeInternal(fs),
	}
}
