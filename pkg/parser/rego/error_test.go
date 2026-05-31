//ff:func feature=policy type=test control=sequence
//ff:what TestRegoHelpers — unit tests for the pure rego parser helper functions
package rego

func (errNotOPA) Error() string { return "plain error" }
