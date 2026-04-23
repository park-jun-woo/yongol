//ff:func feature=cli type=accessor control=sequence
//ff:what usageError.Error — wrapped error message getter
package main

// Error returns the wrapped error message.
func (u *usageError) Error() string { return u.err.Error() }
