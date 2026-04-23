//ff:func feature=migration type=util control=iteration dimension=1
//ff:what ApplyHintsToOps — Operation 리스트에 Hints 반영 (DropTable AllowDestructive / AddColumn Backfill / AlterColumnType Using 등)
package migration

// ApplyHintsToOps returns ops with each operation updated so that its
// AllowDestructive / Backfill / Using fields reflect the supplied
// hints. Call this after Diff so ops' SafetyLevel() respects hints.
func ApplyHintsToOps(ops []Operation, hints *Hints) []Operation {
	if hints == nil {
		return ops
	}
	out := make([]Operation, len(ops))
	for i, op := range ops {
		out[i] = applyHint(op, hints)
	}
	return out
}
