//ff:type feature=cli type=model
//ff:what usageError — cobra usage failure wrapper; main maps it to exit code 2

package main

// usageError wraps any error that should produce exit code 2 (POSIX usage
// error convention). Uses include cobra Args validation failures and
// unknown-flag errors routed via FlagErrorFunc.
type usageError struct{ err error }
