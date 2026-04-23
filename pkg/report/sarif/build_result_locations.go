//ff:func feature=report type=util control=sequence topic=sarif
//ff:what buildResultLocations — d.File 이 있을 때 SARIF Location 배열 구성

package sarif

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// buildResultLocations returns a single-element Location slice for a
// diagnostic carrying a file path. When the file path is empty the caller
// should omit the Locations field entirely (nil is returned).
func buildResultLocations(d diagnostic.Diagnostic, specsDir, absSpecs string) []Location {
	if d.File == "" {
		return nil
	}
	return []Location{{
		PhysicalLocation: PhysicalLocation{
			ArtifactLocation: ArtifactLocation{URI: relativeArtifactURI(d.File, specsDir, absSpecs)},
			Region:           regionOrNil(d.Line),
		},
	}}
}
