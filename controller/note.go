package controller

import (
	"net/http"

	"github.com/blue-jay/blueshift/domain"
	"github.com/blue-jay/blueshift/lib/flight"
	"github.com/blue-jay/blueshift/model/note"

	"github.com/blue-jay/core/pagination"
)

/*var (
	uri = "/notepad"
)*/

// Load the routes.
/*func Load() {
	c := router.Chain(acl.DisallowAnon)
	router.Get(uri, Index, c...)
	router.Get(uri+"/create", Create, c...)
	router.Post(uri+"/create", Store, c...)
	router.Get(uri+"/view/:id", Show, c...)
	router.Get(uri+"/edit/:id", Edit, c...)
	router.Patch(uri+"/edit/:id", Update, c...)
	router.Delete(uri+"/:id", Destroy, c...)
}*/

// NoteHandler represents the services required for this controller.
type NoteHandler struct {
	UserService domain.UserCase
	ViewService domain.ViewCase
	Uri         string
}

// Index displays the items.
func (h *NoteHandler) Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// Create a pagination instance with a max of 10 results.
	p := pagination.New(r, 10)

	items, _, err := note.ByUserIDPaginate(c.DB, c.UserID, p.PerPage, p.Offset)
	if err != nil {
		c.FlashErrorGeneric(err)
		items = []note.Item{}
	}

	count, err := note.ByUserIDCount(c.DB, c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
	}

	// Calculate the number of pages.
	p.CalculatePages(count)

	v := c.View.New("note/index")
	v.Vars["items"] = items
	v.Vars["pagination"] = p
	v.Render(w, r)
}

// Create displays the create form.
func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	v := c.View.New("note/create")
	c.Repopulate(v.Vars, "name")
	v.Render(w, r)
}

// Store handles the create form submission.
func (h *NoteHandler) Store(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	if !c.FormValid("name") {
		h.Create(w, r)
		return
	}

	_, err := note.Create(c.DB, r.FormValue("name"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		h.Create(w, r)
		return
	}

	c.FlashSuccess("Item added.")
	c.Redirect(h.Uri)
}

// Show displays a single item.
func (h *NoteHandler) Show(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	item, _, err := note.ByID(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(h.Uri)
		return
	}

	v := c.View.New("note/show")
	v.Vars["item"] = item
	v.Render(w, r)
}

// Edit displays the edit form.
func (h *NoteHandler) Edit(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	item, _, err := note.ByID(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(h.Uri)
		return
	}

	v := c.View.New("note/edit")
	c.Repopulate(v.Vars, "name")
	v.Vars["item"] = item
	v.Render(w, r)
}

// Update handles the edit form submission.
func (h *NoteHandler) Update(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	if !c.FormValid("name") {
		h.Edit(w, r)
		return
	}

	_, err := note.Update(c.DB, r.FormValue("name"), c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		h.Edit(w, r)
		return
	}

	c.FlashSuccess("Item updated.")
	c.Redirect(h.Uri)
}

// Destroy handles the delete form submission.
func (h *NoteHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	_, err := note.DeleteSoft(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
	} else {
		c.FlashNotice("Item deleted.")
	}

	c.Redirect(h.Uri)
}
