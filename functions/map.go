package functions

import "text/template"

// Map returns a map of functions to use with Go's standard template package.
func Map() template.FuncMap {
	return template.FuncMap{
		"required": required,
	}
}
