package items

import (
	"bytes"
	"net/http"
	"html/template"
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/models/item"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
	"github.com/charmbracelet/log"
)

var itemTpl = template.Must(template.ParseFiles("./templates/item/item.html", "./templates/item/item_overview.html", "./templates/common.html"))

type ItemError struct {
	status int
	wrapped error
}

func NewItemError(status int, err error) *ItemError {
	return &ItemError{ status, err }
}

func (e ItemError) Error() string {
	return e.wrapped.Error()
}

func (e ItemError) Status() int {
	return e.status
}

func ItemList(r *http.Request, backend data.ItemRepo) (bytes.Buffer, *ItemError) {
	log.Debug("Generating Item Overview")
	b := bytes.Buffer{}

	kinds, err := backend.ListItemKinds(r.Context())
	if err != nil {
		log.Error("Error while listing item kinds", "error", err)
		return b, NewItemError(http.StatusInternalServerError, err)
	}

	slots, err := backend.ListItemSlots(r.Context())
	if err != nil {
		log.Error("Error while listing item slots", "error", err)
		return b, NewItemError(http.StatusInternalServerError, err)
	}

	items, err := backend.ListItems(r.Context())
	if err != nil {
		log.Error("Error while listing items", "error", err)
		return b, NewItemError(http.StatusInternalServerError, err)
	}
	log.Debug("Found Items", "count", len(items))

	vm := item.ItemOverviewView{ Kinds: kinds, Slots: slots, Items: items }

	if err := itemTpl.ExecuteTemplate(&b, "item_browser", vm); err != nil {
		log.Error("Error while executing item_browser template", "error", err)
		return b, NewItemError(http.StatusInternalServerError, err)
	}

	return b, nil
}

func ItemComponent(it sqlc.ItemView) (bytes.Buffer, *ItemError) {
	b := bytes.Buffer{}
	if err := itemTpl.ExecuteTemplate(&b, "item_component", it); err != nil {
		log.Error("Error while executing item_component template", "error", err)
		return b, NewItemError(http.StatusInternalServerError, err)
	}

	return b, nil
}

func ItemEditComponent(r *http.Request, backend data.ItemRepo, it *sqlc.ItemView) (bytes.Buffer, *ItemError) {
	b := bytes.Buffer{}

	if it == nil {
		log.Debug("Item to edit was nil. Initializing empty item.")
		it = &sqlc.ItemView{}
	}

	kinds, err := backend.ListItemKinds(r.Context())
	if err != nil {
		log.Error("Error while listing item kinds", "error", err)
		return b, NewItemError(http.StatusInternalServerError, err)
	}

	slots, err := backend.ListItemSlots(r.Context())
	if err != nil {
		log.Error("Error while listing item slots", "error", err)
		return b, NewItemError(http.StatusInternalServerError, err)
	}

	vm := item.ItemEditView { Kinds: kinds, Slots: slots, Item: *it }

	if err := itemTpl.ExecuteTemplate(&b, "item_edit_component", vm); err != nil {
		log.Error("Error while executing item_edit_component template", "error", err)
		return b, NewItemError(http.StatusInternalServerError, err)
	}

	return b, nil
}

// Not enough shared code in creating or updating items to warrent breaking it out here
// func CreateItem(w http.ResponseWriter, r *http.Request, backend data.ItemRepo, it sqlc.ItemView) error {
// }

// func UpdateItem(w http.ResponseWriter, r *http.Request, backend data.ItemRepo, it sqlc.ItemView) error {
// }
