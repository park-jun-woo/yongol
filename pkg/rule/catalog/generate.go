//ff:func feature=rule type=util control=sequence topic=catalog
//ff:what go:generate — copies rulebook.md from the repo root into this package (//go:embed cannot reference parent paths)
package catalog

// Because go:embed cannot reference files outside the package directory (and
// rejects symlinks as irregular files), a verbatim copy is kept here. Sync it
// after editing the root rulebook by running:
//
//	go generate ./pkg/rule/catalog
//
// A missing sync fails TestEmbedInSyncWithRootRulebook (byte-equal guard), so
// drift is caught by the mandatory pre-commit `go test ./...` run.

//go:generate cp ../../../rulebook.md ./rulebook.md
