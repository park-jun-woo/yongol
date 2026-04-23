package boot

// defaultHeaderLimit mirrors Go stdlib http.DefaultMaxHeaderBytes (1 MiB).
// Used when manifest.backend.http.header_limit is unset or unparseable.
const defaultHeaderLimit = int64(1 << 20)
