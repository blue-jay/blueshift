package repository

import (
	"net/http"

	"github.com/blue-jay/blueshift/domain"
	"github.com/blue-jay/core/view"
)

// ViewRepo represents a service for managing routes.
type ViewRepo struct {
	view view.Info
}

// NewViewRepo returns the service for managing routes.
func NewViewRepo(config view.Info) *ViewRepo {
	s := new(ViewRepo)
	s.view = config
	return s
}

// New returns returns the service for managing routes.
func (v ViewRepo) New(templateList ...string) domain.ViewRepo {
	repo := new(ViewRepo)
	repo.view = *v.view.New(templateList...)
	return &v
}

// Render outputs the template to the ResponseWriter.
func (v *ViewRepo) Render(w http.ResponseWriter, r *http.Request) error {
	return v.view.Render(w, r)
}

/*

// SetFolder sets the folder containing the templates.
func (v *ViewRepo) SetFolder(s string) {

}

// SetExtension sets the extensions of the templates.
func (v *ViewRepo) SetExtension(s string) {

}

// SetBaseTemplate sets the base template to render.
func (v *ViewRepo) SetBaseTemplate(s string) {

}

// SetTemplate sets the template to render.
func (v *ViewRepo) SetTemplate(s string) {
	v.view = *v.view.New(s)
}

// AddVar adds a variable to the template variable map.
func (v *ViewRepo) AddVar(key string, value interface{}) {

}

// DelVar removes a variable from the template variable map.
func (v *ViewRepo) DelVar(key string) {

}

// GetVar returns a value from the template variable map.
func (v *ViewRepo) GetVar(key string) interface{} {
	return nil
}

// SetVars sets the template variable map.
func (v *ViewRepo) SetVars(vars domain.ViewVars) {

}*/
