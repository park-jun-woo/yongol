//ff:type feature=tsx-parser type=model
//ff:what FormField — react-hook-form register('name', opts) 선언 1건

package tsx

// FormField captures a react-hook-form register('name', opts) invocation.
type FormField struct {
	Name     string
	Required bool
	Line     int
	Col      int
}
