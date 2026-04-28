//ff:func feature=cli-init type=util control=sequence
//ff:what ValidateProjectID — ensure ProjectID follows supported naming rules

package cliinit

import (
	"fmt"
	"regexp"
)

// projectIDPattern restricts ProjectID to alphanumerics and underscores, with
// the first character being a letter. That matches PascalCase ("MyApp") and
// snake_case ("my_app") inputs. Hyphens / dots / path separators are rejected.
var projectIDPattern = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]*$`)

// ValidateProjectID ensures the CLI positional argument follows the supported
// naming rules. Returns an error message suitable for surfacing to the user.
func ValidateProjectID(id string) error {
	if id == "" {
		return fmt.Errorf("ProjectID is empty")
	}
	if !projectIDPattern.MatchString(id) {
		return fmt.Errorf("ProjectID %q must start with a letter and contain only [A-Za-z0-9_] (PascalCase or snake_case)", id)
	}
	return nil
}
