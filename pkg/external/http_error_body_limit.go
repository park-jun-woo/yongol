package external

// httpErrorBodyLimit is the maximum number of response body bytes preserved
// in an HTTP error message returned by readHTTPSource. Keeps diagnostics
// readable when the remote service returns verbose HTML or large JSON errors.
const httpErrorBodyLimit = 1024
