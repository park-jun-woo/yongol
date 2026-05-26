//ff:func feature=gen-ir type=generator control=sequence
//ff:what BuildInfraPlan -- manifest + prepared.State → InfraPlan 변환

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// BuildInfraPlan inspects manifest + prepared.State to determine which
// infrastructure adapters are active and populates their configuration.
// Pointer fields in the returned plan are nil when the corresponding
// subsystem is inactive.
func BuildInfraPlan(fs *yongol.Fullstack, ps *prepared.State) *InfraPlan {
	plan := &InfraPlan{}

	if ps.ActiveBackends.Session != nil {
		plan.Session = &SessionConfig{
			Backend: ps.ActiveBackends.Session.Backend,
		}
	}

	if ps.ActiveBackends.Cache != nil {
		plan.Cache = &CacheConfig{
			Backend: ps.ActiveBackends.Cache.Backend,
		}
	}

	if ps.ActiveBackends.Queue != nil {
		plan.Queue = &QueueConfig{
			Backend: ps.ActiveBackends.Queue.Backend,
		}
	}

	if ps.Auth.Present {
		plan.Auth = buildAuthInfraConfig(fs, ps)
	}

	return plan
}
