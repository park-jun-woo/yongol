//ff:type feature=cli-init type=model
//ff:what skeletonFile — describes one file to materialize under the target directory

package cliinit

// skeletonFile describes one file to materialize under the target directory.
// `rendered=true` runs the file through text/template with templateData;
// `rendered=false` copies the embedded bytes verbatim (used for files that
// contain no placeholders, like sqlc.yaml).
type skeletonFile struct {
	srcEmbed string // path within the embed.FS
	destRel  string // path relative to opts.Dir
	rendered bool
}
