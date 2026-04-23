//ff:type feature=generate type=model
//ff:what Cache — cache 백엔드 파생 상태 (non-zero 보장)

package prepared

// Cache carries the derived cache backend configuration used by
// codegen. Present only when manifest declares cache.backend OR SSaC
// uses cache.* calls; otherwise State.ActiveBackends.Cache is nil and
// no cache blocks are emitted.
//
// Backend is one of "postgres" or "memory"; never empty.
type Cache struct {
	Backend string
}
