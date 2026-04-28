//ff:func feature=cli-init type=util control=sequence
//ff:what renderTemplate — load a template from the embed.FS and execute against data

package cliinit

import (
	"path/filepath"
	"strings"
	"text/template"
)

// renderTemplate loads a template from the embed.FS and executes it against
// the supplied data. A fresh *template.Template is built per call to keep the
// function easy to reason about; the cost is negligible (N templates, each
// parsed once per `yongol init` invocation).
func renderTemplate(srcEmbed string, data templateData) ([]byte, error) {
	raw, err := templateFiles.ReadFile(srcEmbed)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New(filepath.Base(srcEmbed)).Parse(string(raw))
	if err != nil {
		return nil, err
	}
	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return []byte(buf.String()), nil
}
