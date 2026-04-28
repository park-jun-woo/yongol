//ff:func feature=cli-init type=util control=sequence
//ff:what skeletonDirs — directory tree that must exist regardless of template files

package cliinit

// skeletonDirs returns the directory tree that must exist regardless of
// whether any template file lands in it. Ordering is from shallow to deep so
// os.MkdirAll never has to backfill intermediate nodes in later entries.
func skeletonDirs() []string {
	return []string{
		"specs",
		"specs/api",
		"specs/db",
		"specs/db/queries",
		"specs/service",
		"specs/states",
		"specs/policy",
		"specs/frontend",
		"specs/frontend/pages",
		"specs/frontend/components",
		"specs/tests",
	}
}
