//ff:type feature=generate type=model
//ff:what Queue — queue 백엔드 파생 상태 (non-zero 보장)

package prepared

// Queue carries the derived queue backend configuration used by
// codegen. Present only when manifest declares queue.backend OR SSaC
// uses @publish/@subscribe constructs; otherwise
// State.ActiveBackends.Queue is nil.
//
// Backend is one of "postgres" or "memory"; never empty.
type Queue struct {
	Backend string
}
