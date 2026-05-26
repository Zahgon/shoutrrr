package standard

import (
	"text/template"
)

// Templater is the standard implementation of ApplyTemplate using the "text/template" library
type Templater struct {
	templates map[string]*template.Template
}

// GetTemplate attempts to retrieve the template identified with id
func (templater *Templater) GetTemplate(id string) (template *template.Template, found bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SetTemplateString creates a new template from the body and assigning it the id
func (templater *Templater) SetTemplateString(id string, body string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetTemplateFile creates a new template from the file and assigning it the id
func (templater *Templater) SetTemplateFile(id string, file string) error {
	_ = "STUB: not implemented"
	return nil
}
