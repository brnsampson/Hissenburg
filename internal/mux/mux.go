package mux

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	//"github.com/charmbracelet/log"
	"github.com/brnsampson/Hissenburg/data"
	roll "github.com/brnsampson/Hissenburg/logic/modelroller"
	"github.com/brnsampson/Hissenburg/models/character"
	"github.com/brnsampson/Hissenburg/models/status"
	"github.com/brnsampson/Hissenburg/models/item"
	"github.com/brnsampson/Hissenburg/models/inventory"
)

var assets = http.FileServer(http.Dir("./assets"))
var indexTpl = template.Must(template.ParseFiles("./templates/index.html"))
var charTpl = template.Must(template.ParseFiles("./templates/traits.html", "./templates/status.html", "./templates/inventory.html", "./templates/identity.html", "./templates/character.html", "./templates/list_chars.html"))

func New(storePath string) (*chi.Mux, error) {
	chandler, err := NewCharHandler(storePath)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	router.Get("/", getIndex)
	router.Mount("/character", chandler)
	router.Mount("/assets", http.StripPrefix("/assets", assets))
	return router, nil
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	if err := indexTpl.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type CharHandler struct {
	chi.Router
	backend data.CharBackend
	itemBackend data.ItemBackend
}

func NewCharHandler(storePath string) (*CharHandler, error) {
	backend, err := data.NewCharData(storePath)
	if err != nil {
		return nil, err
	}

	itemBackend, err := data.NewItemData()
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	ch := CharHandler{router, backend, itemBackend}

	router.Get("/", ch.listChars)
	router.Post("/generate", ch.postGenerate)
	router.Route("/{name}/{surname}", func(router chi.Router) {
		router.Get("/", ch.getChar)
		router.Post("/", ch.postChar)
		router.Delete("/", ch.deleteChar)
		router.Get("/identity", ch.getIdentity)
		router.Get("/identity/edit", ch.editIdentity)
		router.Put("/identity", ch.putIdentity)
		router.Get("/status", ch.getStatus)
		router.Get("/status/edit", ch.editStatus)
		router.Put("/status", ch.putStatus)
		router.Get("/traits", ch.getTraits)
		router.Get("/traits/edit", ch.editTraits)
		router.Put("/traits", ch.putTraits)
		router.Get("/inventory", ch.getInventory)
		router.Put("/inventory", ch.putInventory)
		router.Route("/inventory/backpack", func(router chi.Router) {
			router.Put("/", ch.putBackpack)
			router.Get("/edit", ch.editBackpack)
			router.Put("/{slot}", ch.putBackpackSlot)
			router.Delete("/{slot}", ch.deleteBackpack)
		})
		router.Delete("/inventory/ground/{slot}", ch.deleteGround)
	})

	return &ch, nil
}


func (ch *CharHandler) listChars(w http.ResponseWriter, r *http.Request) {
	chars, err := ch.backend.ListCharacters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := charTpl.ExecuteTemplate(w, "character_list", chars); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) postGenerate(w http.ResponseWriter, r *http.Request) {
	char := character.New()
	err := roll.RollChar(ch.backend, ch.itemBackend, &char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = ch.backend.CreateCharacter(&char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s", char.Name, char.Surname), http.StatusSeeOther)
}

func (ch *CharHandler) getChar(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := charTpl.ExecuteTemplate(w, "character", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) postChar(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char := character.New()
	char.Name = name
	char.Surname = surname
	_, err := ch.backend.CreateCharacter(&char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) deleteChar(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	_, err := ch.backend.DeleteCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

    http.Redirect(w, r, "/character", http.StatusSeeOther)
}

func (ch *CharHandler) getIdentity(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := charTpl.ExecuteTemplate(w, "identity", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) editIdentity(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := charTpl.ExecuteTemplate(w, "edit_identity", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) putIdentity(w http.ResponseWriter, r *http.Request) {
	oldName := chi.URLParam(r, "name")
	oldSurname := chi.URLParam(r, "surname")

	name := r.FormValue("name")
	surname := r.FormValue("surname")
	gender := r.FormValue("gender")
	background := r.FormValue("background")
	ageStr := r.FormValue("age")
	age64, err := strconv.ParseUint(ageStr, 10, 16)
	if err != nil {
		http.Error(w, "Age must be an integer between 0 and 65535", http.StatusInternalServerError)
	}
	age := uint16(age64)

	char, err := ch.backend.GetCharacter(oldName, oldSurname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	char.Name = name
	char.Surname = surname
	char.Gender = character.GenderFromString(gender)
	char.Background = background
	char.Age = age
	char, err = ch.backend.UpdateCharacter(oldName, oldSurname, char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/identity", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) getInventory(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := charTpl.ExecuteTemplate(w, "inventory", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) putInventory(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	iStrings := r.Form["inventory"]
	iKindStrings := r.Form["inventoryTypes"]
	if len(iStrings) != len(iKindStrings) {
		http.Error(w, "The list of items and list of item types were of different lengths", http.StatusInternalServerError)
	}

	items := make([]item.Item, 0)
	for i, slot := range iStrings {
		kind := item.TypeFromString(iKindStrings[i])
		invItem, err := ch.itemBackend.GetItem(kind, slot)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		items = append(items, invItem)
	}

	inv, err := inventory.Unpack(items)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	char.Inventory = inv

	char, err = ch.backend.UpdateCharacter(name, surname, char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) putBackpack(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	iStrings := r.Form["item_name"]
	iKindStrings := r.Form["item_type"]
	if len(iStrings) != len(iKindStrings) {
		http.Error(w, "The list of items and list of item types were of different lengths", http.StatusInternalServerError)
	}

	items := make([]item.Item, 0)
	for i, slot := range iStrings {
		kind := item.TypeFromString(iKindStrings[i])
		invItem, err := ch.itemBackend.GetItem(kind, slot)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		items = append(items, invItem)
	}

	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for i, it := range items {
		_, err := char.Inventory.SetBackpack(i, it)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	char, err = ch.backend.UpdateCharacter(name, surname, char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) putBackpackSlot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	slot, err := strconv.Atoi(chi.URLParam(r, "slot"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	iName := r.Form["item_name"][slot]
	iType := r.Form["item_type"][slot]

	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	kind := item.TypeFromString(iType)
	i, err := ch.itemBackend.GetItem(kind, iName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = char.Inventory.SetBackpack(slot, i)

	_, err = ch.backend.UpdateCharacter(name, surname, char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) editBackpack(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tl := item.ListItemTypes()
	vm := char.Inventory.IntoView().IntoEditView(char.Name, char.Surname, tl)

	if err := charTpl.ExecuteTemplate(w, "edit_backpack", vm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func (ch *CharHandler) deleteBackpack(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	slot, err := strconv.Atoi(chi.URLParam(r, "slot"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	old, err := char.Inventory.SetBackpack(slot, item.Empty())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	char.Inventory.AddToGround(old)

	char, err = ch.backend.UpdateCharacter(name, surname, char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) deleteGround(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	slot, err := strconv.Atoi(chi.URLParam(r, "slot"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if slot < 0 || slot >= len(char.Inventory.Ground) {
		http.Error(w, "Addempted to delete a ground item out of range", http.StatusInternalServerError)
	}
	char.Inventory.Ground = append(char.Inventory.Ground[:slot], char.Inventory.Ground[slot + 1:]...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	char, err = ch.backend.UpdateCharacter(name, surname, char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) getStatus(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := charTpl.ExecuteTemplate(w, "status", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) editStatus(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := charTpl.ExecuteTemplate(w, "edit_status", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) putStatus(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	hpStr := r.FormValue("hp")
	hp64, err := strconv.ParseUint(hpStr, 10, 8)
	if err != nil {
		http.Error(w, "HP must be an integer between 0 and 256 and less than MaxStr", http.StatusInternalServerError)
	}
	hp := uint8(hp64)

	strStr := r.FormValue("str")
	str64, err := strconv.ParseUint(strStr, 10, 8)
	if err != nil {
		http.Error(w, "Str must be an integer between 0 and 256 and less than MaxStr", http.StatusInternalServerError)
	}
	str := uint8(str64)

	dexStr := r.FormValue("dex")
	dex64, err := strconv.ParseUint(dexStr, 10, 8)
	if err != nil {
		http.Error(w, "Dex must be an integer between 0 and 256 and less than MaxDex", http.StatusInternalServerError)
	}
	dex := uint8(dex64)

	willStr := r.FormValue("will")
	will64, err := strconv.ParseUint(willStr, 10, 8)
	if err != nil {
		http.Error(w, "Will must be an integer between 0 and 256 and less than MaxWill", http.StatusInternalServerError)
	}
	will := uint8(will64)

	maxhpStr := r.FormValue("maxhp")
	maxhp64, err := strconv.ParseUint(maxhpStr, 10, 8)
	if err != nil {
		http.Error(w, "MaxHP must be an integer between 0 and 256", http.StatusInternalServerError)
	}
	maxhp := uint8(maxhp64)

	maxstrStr := r.FormValue("maxstr")
	maxstr64, err := strconv.ParseUint(maxstrStr, 10, 8)
	if err != nil {
		http.Error(w, "Str must be an integer between 0 and 256", http.StatusInternalServerError)
	}
	maxstr := uint8(maxstr64)

	maxdexStr := r.FormValue("maxdex")
	maxdex64, err := strconv.ParseUint(maxdexStr, 10, 8)
	if err != nil {
		http.Error(w, "Dex must be an integer between 0 and 256", http.StatusInternalServerError)
	}
	maxdex := uint8(maxdex64)

	maxwillStr := r.FormValue("maxwill")
	maxwill64, err := strconv.ParseUint(maxwillStr, 10, 8)
	if err != nil {
		http.Error(w, "Will must be an integer between 0 and 256", http.StatusInternalServerError)
	}
	maxwill := uint8(maxwill64)

	status := status.New()
	status.HP = hp
	status.Str = str
	status.Dex = dex
	status.Will = will
	status.MaxHP = maxhp
	status.MaxStr = maxstr
	status.MaxDex = maxdex
	status.MaxWill = maxwill

	err = ch.backend.UpdateStatus(name, surname, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/status", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) getTraits(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := charTpl.ExecuteTemplate(w, "traits", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) editTraits(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := charTpl.ExecuteTemplate(w, "edit_traits", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) putTraits(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	physique := r.FormValue("physique")
	skin := r.FormValue("skin")
	hair := r.FormValue("hair")
	face := r.FormValue("face")
	speech := r.FormValue("speech")
	clothing := r.FormValue("clothing")
	virtue := r.FormValue("virtue")
	vice := r.FormValue("vice")
	reputation := r.FormValue("reputation")
	misfortune := r.FormValue("misfortune")


	traits := character.NewTraits()
	traits.Physique = physique
	traits.Skin = skin
	traits.Hair = hair
	traits.Face = face
	traits.Speech = speech
	traits.Clothing = clothing
	traits.Virtue = virtue
	traits.Vice = vice
	traits.Reputation = reputation
	traits.Misfortune = misfortune

	char, err := ch.backend.GetCharacter(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	char.Traits = traits

	char, err = ch.backend.UpdateCharacter(name, surname, char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/traits", name, surname), http.StatusSeeOther)
}
