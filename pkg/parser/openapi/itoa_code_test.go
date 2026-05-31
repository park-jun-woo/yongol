//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestOpenAPIHelpers — unit tests for the pure openapi parser helper functions
package openapi

func itoaCode(c int) string {
	// small local itoa for 3-digit status codes
	return string(rune('0'+c/100)) + string(rune('0'+(c/10)%10)) + string(rune('0'+c%10))
}
