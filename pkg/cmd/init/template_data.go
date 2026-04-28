//ff:type feature=cli-init type=model
//ff:what templateData — context passed to text/template when rendering skeleton files

package cliinit

// templateData is the context passed to text/template when rendering
// skeleton files. Keep fields exported so template authors can reference
// them via `{{.ProjectID}}` and friends.
type templateData struct {
	ProjectID           string
	ProjectIDNormalized string
	Description         string
	Module              string
}
