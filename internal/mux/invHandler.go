package mux

import (
	"strconv"
	"github.com/go-chi/chi/v5"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/models/inventory"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
	"github.com/charmbracelet/log"
)

var invTpl = template.Must(template.ParseFiles("./templates/inv/inv.html", "./templates/common.html"))

type InvHandler struct {
	chi.Router
	backend data.InventoryRepo
}

func NewInvHandler() (*InvHandler, error) {
	db, err := sql.Open("sqlite3", "file:./db/hiss.sqlite3")
	if err != nil {
		return nil, err
	}

	backend := data.NewRepo(db)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	ih := InvHandler{router, backend}

	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", ih.getInv)
		router.Get("/edit", ih.editInv)
		router.Put("/", ih.putInv)
	})

	return &ih, nil
}

func (ih *InvHandler) getInv(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("Could not parse id of inventory as integer", "error", err, "id", idStr)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	it, err := ih.backend.GetInventory(r.Context(), id)
	if err != nil {
		log.Error("Error while looking up inventory", "inventory", id, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	if err := itemTpl.ExecuteTemplate(w, "inventory_component", it); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ih *InvHandler) editInv(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("Could not parse id of inventory as integer", "error", err, "id", idStr)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	it, err := ih.backend.GetInventory(r.Context(), id)
	if err != nil {
		log.Error("Error while looking up inventory", "inventory", id, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	if err := itemTpl.ExecuteTemplate(w, "inventory_edit", it); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ih *InvHandler) putInv(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("Could not parse id of inventory as integer", "error", err, "id", idStr)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	headStr := r.FormValue("head")
	head, err := strconv.ParseInt(headStr, 10, 64)
	if err != nil {
		log.Error("Could not parse id of head item for inventory as integer", "error", err, "inventory", id, "head", headStr)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	torsoStr := r.FormValue("torso")
	torso, err := strconv.ParseInt(torsoStr, 10, 64)
	if err != nil {
		log.Error("Could not parse id of torso item for inventory as integer", "error", err, "inventory", id, "torso", torsoStr)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	left_handStr := r.FormValue("left_hand")
	left_hand, err := strconv.ParseInt(left_handStr, 10, 64)
	if err != nil {
		log.Error("Could not parse id of left_hand item for inventory as integer", "error", err, "inventory", id, "left_hand", left_handStr)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	right_handStr := r.FormValue("right_hand")
	right_hand, err := strconv.ParseInt(right_handStr, 10, 64)
	if err != nil {
		log.Error("Could not parse id of right_hand item for inventory as integer", "error", err, "inventory", id, "right_hand", right_handStr)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	backpackStrs := r.PostForm["backpack"]
	backpack := make([]int64, len(backpackStrs), len(backpackStrs))
	for i, iidStr := range backpackStrs {
		itid, err := strconv.ParseInt(iidStr, 10, 64)
		if err != nil {
			log.Error("Could not parse id of backpack item for inventory as integer", "error", err, "inventory", id, "backpack", iidStr, "index", i)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		backpack[i] = itid
	}

	bonusStrs := r.PostForm["bonus"]
	bonus := make([]int64, len(bonusStrs), len(bonusStrs))
	for i, iidStr := range bonusStrs {
		itid, err := strconv.ParseInt(iidStr, 10, 64)
		if err != nil {
			log.Error("Could not parse id of bonus space item for inventory as integer", "error", err, "inventory", id, "bonus", iidStr, "index", i)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bonus[i] = itid
	}

	groundStrs := r.PostForm["ground"]
	ground := make([]int64, len(groundStrs), len(groundStrs))
	for i, iidStr := range groundStrs {
		itid, err := strconv.ParseInt(iidStr, 10, 64)
		if err != nil {
			log.Error("Could not parse id of ground item for inventory as integer", "error", err, "inventory", id, "ground", iidStr, "index", i)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ground[i] = itid
	}

	inv := inventory.New()
	inv.ID = id
	inv.Equipment.Head = &sqlc.ItemView{ID: head}
	inv.Equipment.Torso = &sqlc.ItemView{ID: torso}
	inv.Equipment.LeftHand = &sqlc.ItemView{ID: left_hand}
	inv.Equipment.RightHand = &sqlc.ItemView{ID: right_hand}

	for _, bid := range backpack {
		inv.AddToBackpack(sqlc.ItemView{ID: bid}, 6)
	}

	for _, bid := range bonus {
		inv.AddToBonusSpace(sqlc.ItemView{ID: bid})
	}

	for _, bid := range ground {
		inv.AddToGround(sqlc.ItemView{ID: bid})
	}

	err = ih.backend.UpdateInventory(r.Context(), inv)
	if err != nil {
		log.Error("Error while updating inventory", "inventory", inv.ID, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}

    http.Redirect(w, r, "/inventory/" + idStr, http.StatusSeeOther)
}
