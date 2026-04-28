package infra

// packageEmitters dispatches from `interface.yaml` package name to the
// concrete emitter. Packages that declare no Go-interface mismatch (none at
// time of Phase002 write — every DB-using ssac package currently needs at
// least one glue conversion) could fall back to emitPostgresImplGeneric.
var packageEmitters = map[string]packageEmitter{
	"cache":   emitCacheWrapper,
	"session": emitSessionWrapper,
	"queue":   emitQueueWrapper,
	"auth":    emitAuthWrapper,
}
