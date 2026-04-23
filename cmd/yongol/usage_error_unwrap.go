//ff:func feature=cli type=accessor control=sequence
//ff:what usageError.Unwrap — exposes underlying error for errors.Is / errors.As
package main

// Unwrap exposes the underlying error for errors.Is / errors.As.
func (u *usageError) Unwrap() error { return u.err }
