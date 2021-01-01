package template

import (
	"io"
)

// Render handles the actual rendering of the templates
func (t *Template) Render(
	w io.Writer,
	name string,
	data interface{},
) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
