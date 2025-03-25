package mux

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/go-chi/chi/v5"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
	"github.com/brnsampson/Hissenburg/internal/routes/items"
	"github.com/charmbracelet/log"
)

var itemTpl = template.Must(template.ParseFiles("./templates/item/item.html", "./templates/item/item_overview.html", "./templates/common.html"))

type ItemHandler struct {
	chi.Router
	backend data.ItemRepo
}

func NewItemHandler() (*ItemHandler, error) {
	db, err := sql.Open("sqlite3", "file:./db/hiss.sqlite3")
	if err != nil {
		return nil, err
	}

	backend := data.NewRepo(db)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	ih := ItemHandler{router, backend}

	router.Get("/", ih.viewOverview)
	router.Route("/{kind}/{name}", func(router chi.Router) {
		router.Get("/", ih.getComponentKN)
		router.Get("/view", ih.viewItem)
		router.Get("/edit", ih.getEditComponentKN)
		router.Post("/", ih.postItem)
		router.Put("/", ih.putItem)
	})
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", ih.getComponent)
		router.Get("/edit", ih.getEditComponent)
	})

	return &ih, nil
}

func (ih *ItemHandler) viewOverview(w http.ResponseWriter, r *http.Request) {
	log.Debug("Generating Item Overview")

	b, err := items.ItemList(r, ih.backend)
	if err != nil {
		log.Error("Error while generating item list view", "error", err)
		http.Error(w, err.Error(), err.Status())
	}

	ReturnPageView(w, r, &b)
}

func (ih *ItemHandler) viewItem(w http.ResponseWriter, r *http.Request) {
	kind := strings.ToLower(chi.URLParam(r, "kind"))
	name := strings.ToLower(chi.URLParam(r, "name"))
	it, err := ih.backend.GetItemFromKindAndName(r.Context(), kind, name)
	if err != nil {
		log.Error("Error while looking up item", "item", name, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	b, cErr := items.ItemComponent(it)
	if cErr != nil {
		http.Error(w, cErr.Error(), cErr.Status())
	}

	ReturnPageView(w, r, &b)
}

func (ih *ItemHandler) getComponent(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("Error while parsing item id", "id", idStr, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	it, err := ih.backend.GetItemView(r.Context(), id)
	if err != nil {
		log.Error("Error while looking up item", "id", id, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	b, cErr := items.ItemComponent(it)
	if cErr != nil {
		log.Error("Error while generating item component", "id", id, "item", it.Name, "error", err)
		http.Error(w, cErr.Error(), cErr.Status())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	b.WriteTo(w)
}

func (ih *ItemHandler) getComponentKN(w http.ResponseWriter, r *http.Request) {
	kind := strings.ToLower(chi.URLParam(r, "kind"))
	name := strings.ToLower(chi.URLParam(r, "name"))
	it, err := ih.backend.GetItemFromKindAndName(r.Context(), kind, name)
	if err != nil {
		log.Error("Error while looking up item", "item", name, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	b, cErr := items.ItemComponent(it)
	if cErr != nil {
		log.Error("Error while generating item component", "item", name, "error", err)
		http.Error(w, cErr.Error(), cErr.Status())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	b.WriteTo(w)
}

func (ih *ItemHandler) getEditComponent(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("Error while parsing item id", "id", idStr, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	it, err := ih.backend.GetItemView(r.Context(), id)
	if err != nil {
		log.Error("Error while looking up item", "id", id, "error", err)
		it = sqlc.ItemView{}
	} else {
		log.Debug("Found Item", "name", it.Name, "kind", it.Kind)
	}

	b, cErr := items.ItemEditComponent(r, ih.backend, &it)
	if cErr != nil {
		log.Error("Error while generating item edit component", "id", id, "item", it.Name, "error", err)
		http.Error(w, cErr.Error(), cErr.Status())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	b.WriteTo(w)
}

func (ih *ItemHandler) getEditComponentKN(w http.ResponseWriter, r *http.Request) {
	kind := chi.URLParam(r, "kind")
	name := chi.URLParam(r, "name")
	it, err := ih.backend.GetItemFromKindAndName(r.Context(), kind, name)
	if err != nil {
		log.Debug("Item to edit not found", "item", name)
		it = sqlc.ItemView{}
	} else {
		log.Debug("Found Item", "name", it.Name, "kind", it.Kind)
	}

	b, cErr := items.ItemEditComponent(r, ih.backend, &it)
	if cErr != nil {
		log.Error("Error while generating item edit component", "item", name, "error", err)
		http.Error(w, cErr.Error(), cErr.Status())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	b.WriteTo(w)
}

func (ih *ItemHandler) postItem(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	kind := r.FormValue("kind")
	slot := r.FormValue("slot")
	if name == "" {
		log.Error("Error while parsing new item", "field", "name", "error", "Required field is empty")
		http.Error(w, "Name field is blank", http.StatusInternalServerError)
		return
	}
	if kind == "" {
		log.Error("Error while parsing new item", "field", "kind", "error", "Required field is empty")
		http.Error(w, "Kind field is blank", http.StatusInternalServerError)
		return
	}
	if slot == "" {
		log.Error("Error while parsing new item", "field", "slot", "error", "Required field is empty")
		http.Error(w, "Slot field is blank", http.StatusInternalServerError)
		return
	}
	description := r.FormValue("description")
	value, err := strconv.ParseInt(r.FormValue("value"), 10, 64)
	if err != nil {
		log.Error("Error while parsing new item", "field", "value", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dice_count, err := strconv.ParseInt(r.FormValue("dice_count"), 10, 64)
	if err != nil {
		log.Error("Error while parsing new item", "field", "dice_count", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dice_sides, err := strconv.ParseInt(r.FormValue("dice_sides"), 10, 64)
	if err != nil {
		log.Error("Error while parsing new item", "field", "dice_sides", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	armor, err := strconv.ParseInt(r.FormValue("armor"), 10, 64)
	if err != nil {
		log.Error("Error while parsing new item", "field", "armor", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	storage, err := strconv.ParseInt(r.FormValue("storage"), 10, 64)
	if err != nil {
		log.Error("Error while parsing new item", "field", "storage", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	size, err := strconv.ParseInt(r.FormValue("size"), 10, 64)
	if err != nil {
		log.Error("Error while parsing new item", "field", "size", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	active_size, err := strconv.ParseInt(r.FormValue("active_size"), 10, 64)
	if err != nil {
		log.Error("Error while parsing new item", "field", "active_size", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stackable, err := strconv.ParseBool(r.FormValue("stackable"))
	if err != nil {
		log.Error("Error while parsing new item", "field", "stackable", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// I'm not sure what to do with icons yet. Is this more of a thing we give an upload option for?
	//icon := r.FormValue("icon")

	it := sqlc.ItemView{
		ID: -1,
		Name: name,
		Kind: kind,
		Slot: slot,
		Description: description,
		Value: value,
		DiceCount: dice_count,
		DiceSides: dice_sides,
		Armor: armor,
		Storage: storage,
		Size: size,
		ActiveSize: active_size,
		Stackable: stackable,
		Icon: "",
	}


	_, err = ih.backend.CreateItem(r.Context(), it)
	if err != nil {
		log.Error("Error while creating new item in DB", "item", it, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/item/%s/%s", kind, name), http.StatusSeeOther)
}

func (ih *ItemHandler) putItem(w http.ResponseWriter, r *http.Request) {
	kind := strings.ToLower(chi.URLParam(r, "kind"))
	name := chi.URLParam(r, "name")
	// get new values from form
	newName := r.FormValue("name")
	newKind := r.FormValue("kind")
	slot := r.FormValue("slot")
	description := r.FormValue("description")
	value, err := strconv.ParseInt(r.FormValue("value"), 10, 64)
	if err != nil {
		log.Error("Error while parsing item update", "field", "value", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dice_count, err := strconv.ParseInt(r.FormValue("dice_count"), 10, 64)
	if err != nil {
		log.Error("Error while parsing item update", "field", "dice_count", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dice_sides, err := strconv.ParseInt(r.FormValue("dice_sides"), 10, 64)
	if err != nil {
		log.Error("Error while parsing item update", "field", "dice_sides", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	armor, err := strconv.ParseInt(r.FormValue("armor"), 10, 64)
	if err != nil {
		log.Error("Error while parsing item update", "field", "armor", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	storage, err := strconv.ParseInt(r.FormValue("storage"), 10, 64)
	if err != nil {
		log.Error("Error while parsing item update", "field", "storage", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	size, err := strconv.ParseInt(r.FormValue("size"), 10, 64)
	if err != nil {
		log.Error("Error while parsing item update", "field", "size", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	active_size, err := strconv.ParseInt(r.FormValue("active_size"), 10, 64)
	if err != nil {
		log.Error("Error while parsing item update", "field", "active_size", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stackable, err := strconv.ParseBool(r.FormValue("stackable"))
	if err != nil {
		log.Error("Error while parsing item update", "field", "stackable", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// I'm not sure what to do with icons yet. Is this more of a thing we give an upload option for?
	//icon := r.FormValue("icon")

	it, err := ih.backend.GetItemFromKindAndName(r.Context(), kind, name)
	if err != nil {
		log.Error("Error while querying item for update", "kind", kind, "name", name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newName != "" {
		it.Name = newName
	}
	if newKind != "" {
		it.Kind = newKind
	}
	if slot != "" {
		it.Slot = slot
	}
	it.Description  = description
	it.Value = value
	it.DiceCount = dice_count
	it.DiceSides = dice_sides
	it.Armor = armor
	it.Storage = storage
	it.Size = size
	it.ActiveSize = active_size
	it.Stackable = stackable

	err = ih.backend.UpdateItem(r.Context(), it)
	if err != nil {
		log.Error("Error while applying update to DB for item", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/item/%s/%s", newKind, newName), http.StatusSeeOther)
}
