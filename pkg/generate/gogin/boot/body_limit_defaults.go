package boot

// Body-limit defaults emitted into generated main.go when manifest provides
// no body_limit / multipart_limit values. const-only var file so filefunc
// does not require an //ff:func annotation.

const (
	defaultBodyLimit      = int64(1 << 20)  // 1 MiB
	defaultMultipartLimit = int64(32 << 20) // 32 MiB
)
