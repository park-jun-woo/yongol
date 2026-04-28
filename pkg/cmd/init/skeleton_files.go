//ff:func feature=cli-init type=util control=sequence
//ff:what skeletonFiles — files to materialize under the target directory

package cliinit

func skeletonFiles() []skeletonFile {
	return []skeletonFile{
		{"templates/manifest.yaml.tmpl", "specs/manifest.yaml", true},
		{"templates/openapi.yaml.tmpl", "specs/api/openapi.yaml", true},
		{"templates/sqlc.yaml", "specs/db/sqlc.yaml", false},
		{"templates/authz.rego.tmpl", "specs/policy/authz.rego", true},
		{"templates/README.md.tmpl", "README.md", true},
		// gitignore is stored without a leading dot so `go:embed` does not
		// silently skip it (embed ignores files whose name begins with "." or
		// "_"). We rename on write.
		{"templates/gitignore", ".gitignore", false},
	}
}
