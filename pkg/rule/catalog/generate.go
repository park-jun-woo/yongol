//ff:func feature=rule type=util control=sequence topic=catalog
//ff:what go:generate — copies rulebook.md from the repo root into this package (//go:embed cannot reference parent paths)
package catalog

// Because go:embed cannot reference files outside the package directory, sync the copy
// before a release or after editing the rulebook by running:
//
//	go generate ./pkg/rule/catalog
//
// The design surfaces a missing sync at the point where cmd/yongol/main.go calls MustLoad(),
// so any out-of-date copy is caught immediately during functional testing.

//go:generate cp ../../../rulebook.md ./rulebook.md
