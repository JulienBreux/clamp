package functions

import "text/template"

// Map returns a map of functions to use with Go's standard template package.
func Map() template.FuncMap {
	fm := make(template.FuncMap)

	return fm
}
