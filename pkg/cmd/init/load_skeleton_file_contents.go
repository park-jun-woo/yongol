//ff:func feature=cli-init type=util control=selection
//ff:what loadSkeletonFileContents — reads raw or rendered skeleton file bytes

package cliinit

import "fmt"

// loadSkeletonFileContents returns the final byte payload for a skeleton
// file, rendering through text/template when f.rendered is true and copying
// the embed.FS bytes verbatim otherwise.
func loadSkeletonFileContents(data templateData, f skeletonFile) ([]byte, error) {
	switch {
	case f.rendered:
		rendered, err := renderTemplate(f.srcEmbed, data)
		if err != nil {
			return nil, fmt.Errorf("render %s: %w", f.srcEmbed, err)
		}
		return rendered, nil
	default:
		raw, err := templateFiles.ReadFile(f.srcEmbed)
		if err != nil {
			return nil, fmt.Errorf("read embedded %s: %w", f.srcEmbed, err)
		}
		return raw, nil
	}
}
