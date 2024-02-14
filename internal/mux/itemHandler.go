package mux

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"github.com/brnsampson/Hissenburg/data"
)

type ItemHandler struct {
	chi.Router
	backend data.ItemBackend
}

func NewItemHandler(storePath string) (*ItemHandler, error) {
	backend, err := data.NewItemData()
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	ih := ItemHandler{router, backend}

	router.Get("/", ih.listItems)
	router.Get("/types", ih.listItemTypes)
	// So far items are all static data so no post... 😅
	//router.Post("/", ch.postItem)
	router.Route("/{kind}", func(router chi.Router) {
		router.Get("/", ih.listItemsByType)
		router.Route("/{name}", func(router chi.Router) {
			router.Get("/", ih.getItem)
		})
	})

	return &ih, nil
}

func (ih *ItemHandler) listItems(w http.ResponseWriter, r *http.Request) {

}

func (ih *ItemHandler) listItemTypes(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (ih *ItemHandler) listItemsByType(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (ih *ItemHandler) getItem(w http.ResponseWriter, r *http.Request) {
	// TODO
}

