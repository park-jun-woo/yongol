//ff:type feature=generate type=model
//ff:what File — file 스토리지 백엔드 파생 상태 (non-zero 보장)

package prepared

// File carries the derived file storage configuration used by codegen.
// Present only when manifest declares file.backend OR SSaC uses file.*
// calls; otherwise State.ActiveBackends.File is nil.
//
// Backend is one of "local" or "s3"; never empty.
type File struct {
	Backend string
}
