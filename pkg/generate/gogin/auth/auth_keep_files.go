package auth

// authKeepFiles lists the files generateAuth is expected to produce. Any
// other *.go in authDir is a leftover from a previous yongol version and
// will be removed by cleanStaleAuthFiles.
var authKeepFiles = map[string]bool{
	"claim.go":    true,
	"reexport.go": true,
}
