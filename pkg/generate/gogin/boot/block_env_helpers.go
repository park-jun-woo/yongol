//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockEnvHelpers — envInt / envDuration / envStringList / envBool 헬퍼 (top-level)

package boot

// blockEnvHelpers produces top-level function declarations (envInt,
// envDuration, envStringList, envBool) appended after main(). Used by
// DB pool config, CORS config, and any future env-driven block. The
// helpers silently fall back to the provided default on parse failure.
//
// Lines is empty — Funcs slot carries the entire payload. Imports include
// strconv / time / strings / os because main() itself doesn't need them
// just for helper usage; dedup merges with other blocks that already
// import os etc.
func blockEnvHelpers() MainBlock {
	return MainBlock{
		Name:    "env-helpers",
		Imports: envHelperImports,
		Funcs:   envHelperFuncs,
	}
}
