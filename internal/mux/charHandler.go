package mux

import (
	"fmt"
	"bytes"
	"strconv"
	"github.com/go-chi/chi/v5"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/models/character"
	"github.com/brnsampson/Hissenburg/logic/modelroller"
	"github.com/charmbracelet/log"
)

var charTpl = template.Must(template.ParseFiles("./templates/character/character.html", "./templates/character/list.html", "./templates/common.html"))

type CharHandler struct {
	chi.Router
	backend data.CharRepo
	invBackend data.InventoryRepo
}

func NewCharHandler() (*CharHandler, error) {
	db, err := sql.Open("sqlite3", "file:./db/hiss.sqlite3")
	if err != nil {
		return nil, err
	}

	backend := data.NewRepo(db)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	ch := CharHandler{router, backend, backend}

	router.Get("/", ch.listChar)
	router.Get("/new", ch.newChar)
	router.Post("/generate", ch.generateChar)
	router.Post("/", ch.postChar)
	router.Route("/{name}/{surname}", func(router chi.Router) {
		router.Get("/", ch.viewChar)
		router.Put("/", ch.putChar)
		router.Get("/edit", ch.editChar)
		router.Route("/associations", func(router chi.Router) {
			router.Get("/", ch.getAssociations)
			router.Get("/edit", ch.editAssociations)
			router.Put("/", ch.putAssociations)
		})
		router.Route("/identity", func(router chi.Router) {
			router.Get("/", ch.getIdentity)
			router.Get("/edit", ch.editIdentity)
			router.Put("/", ch.putIdentity)
		})
		router.Route("/traits", func(router chi.Router) {
			router.Get("/", ch.getTraits)
			router.Get("/edit", ch.editTraits)
			router.Put("/", ch.putTraits)
		})
		router.Route("/status", func(router chi.Router) {
			router.Get("/", ch.getStatus)
			router.Get("/edit", ch.editStatus)
			router.Put("/", ch.putStatus)
			router.Put("/max", ch.putMaxStatus)
		})
	})

	return &ch, nil
}

func (ch *CharHandler) listChar(w http.ResponseWriter, r *http.Request) {
	uc, err := r.Cookie("username")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Debug("Request rejected: No username cookie present")
			LoginRedirect(w, r)
			return
		}
		// For any other type of error, return a bad request status
		log.Debugf("Request rejected: Error reading auth cookie: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if uc.Valid() != nil {
		log.Debug("Request rejected: username cookie expired or invalid")
		LoginRedirect(w, r)
		return
	}

	log.Debug("Generating Character List")
	chars, err := ch.backend.ListCharacterViews(r.Context())
	if err != nil {
		log.Error("Error while listing characters", "error", err)
		// Technically we might get this error if there were no characters to list...
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Debug("Found characters", "count", len(chars))

	vm := character.CharacterListView{ Characters: chars }

	b := bytes.Buffer{}
	if err := charTpl.ExecuteTemplate(&b, "character_browser", vm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ReturnPageViewUserRequired(w, r, &b)
}

func (ch *CharHandler) generateChar(w http.ResponseWriter, r *http.Request) {
	uc, err := r.Cookie("username")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Debug("Request rejected: No username cookie present")
			LoginRedirect(w, r)
			return
		}
		// For any other type of error, return a bad request status
		log.Debugf("Request rejected: Error reading auth cookie: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if uc.Valid() != nil {
		log.Debug("Request rejected: username cookie expired or invalid")
		LoginRedirect(w, r)
		return
	}
	uname := uc.Value

	pname := r.FormValue("party")
	vname := r.FormValue("village")
	// Create new user (if needed)
	user, err := ch.backend.GetUserFromName(r.Context(), uname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired user", "village", vname, "user", user.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	party, err := ch.backend.GetPartyFromName(r.Context(), pname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired party", "party", vname, "user", user.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create new village (if needed)
	village, err := ch.backend.GetVillageFromName(r.Context(), vname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired village", "village", vname, "user", user.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	assoc := character.Associations{
		CharacterID: -1,
		User: user,
		Party: party,
		Village: village,
	}
	char := character.Character{
		ID: -1,
		Associations: assoc,
	}
	char, err = modelroller.RollCharacter(r.Context(), ch.backend, char)
	if err != nil {
		log.Error("Failed to roll new character (before DB insertion)")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create character in DB
	char, err = ch.backend.CreateCharacter(r.Context(), char)
	if err != nil {
		log.Error("Failed to insert new character into the DB")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Now roll a new inventory...
	inv, err := modelroller.RollInventory(r.Context(), ch.backend)
	inv.ID = char.Inventory
	// update the inventory in the DB
	err = ch.backend.UpdateInventory(r.Context(), inv)
	if err != nil {
		log.Warn("Failed to update inventory for new character", "cid", char.ID, "name", char.Identity.Name, "surname", char.Identity.Surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s", char.Identity.Name, char.Identity.Surname), http.StatusSeeOther)
}

func (ch *CharHandler) postChar(w http.ResponseWriter, r *http.Request) {
	uc, err := r.Cookie("username")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Debug("Request rejected: No username cookie present")
			LoginRedirect(w, r)
			return
		}
		// For any other type of error, return a bad request status
		log.Debugf("Request rejected: Error reading auth cookie: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if uc.Valid() != nil {
		log.Debug("Request rejected: username cookie expired or invalid")
		LoginRedirect(w, r)
		return
	}
	uname := uc.Value

	name := r.FormValue("name")
	surname := r.FormValue("surname")
	vname := r.FormValue("village")
	pname := r.FormValue("party")
	user, err := ch.backend.GetUserFromName(r.Context(), uname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired user", "name", name, "surname", surname, "village", vname, "user", user.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	party, err := ch.backend.GetPartyFromName(r.Context(), pname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired party", "name", name, "surname", surname, "party", vname, "user", user.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	village, err := ch.backend.GetVillageFromName(r.Context(), vname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired village", "name", name, "surname", surname, "village", vname, "user", user.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create new identity with at least name and surname
	id := character.Identity{ Name: name, Surname: surname }
	// Create new status
	status := character.Status{}
	// Create new traits
	traits := character.Traits{}
	// Create new empty inventory
	inv, err := ch.backend.CreateInventory(r.Context())
	if err != nil {
		log.Error("Failed to create new inventory for new character", "name", name, "surname", surname, "village", vname, "user", user.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create new character

	char := character.New()
	char.Associations.User = user
	char.Associations.Village = village
	char.Associations.Party = party
	char.Identity = id
	char.Traits = traits
	char.Status = status
	char.Inventory = inv
	c, err := ch.backend.CreateCharacter(r.Context(), char)
	if err != nil {
		log.Error("Failed to start new character", "name", name, "surname", surname, "user", user, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s", c.Identity.Name, c.Identity.Surname), http.StatusSeeOther)
}

func (ch *CharHandler) newChar(w http.ResponseWriter, r *http.Request) {
	char := character.New()

	b := bytes.Buffer{}
	if err := charTpl.ExecuteTemplate(&b, "character_creator", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ReturnPageViewUserRequired(w, r, &b)
}

func (ch *CharHandler) viewChar(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	log.Debug("Lookind up Character", "name", name, "surname", surname)
	char, err := ch.backend.GetCharacterViewFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Error while looking up character", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Debug("Found character", "character", char)

	b := bytes.Buffer{}
	if err := charTpl.ExecuteTemplate(&b, "character_sheet", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ReturnPageViewUserRequired(w, r, &b)
}

func (ch *CharHandler) editChar(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	log.Debug("Lookind up Character", "name", name, "surname", surname)
	char, err := ch.backend.GetCharacterFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Error while looking up character", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	log.Debug("Found character", "character", char)

	if err := charTpl.ExecuteTemplate(w, "character_edit", char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) putChar(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	char, err := ch.backend.GetCharacterFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Could not find character", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	uname := r.FormValue("user")
	pname := r.FormValue("party")
	vname := r.FormValue("village")
	user, err := ch.backend.GetUserFromName(r.Context(), uname)
	if err != nil {
		log.Error("Could not find user when updating character", "error", err, "name", name, "surname", surname, "user", uname)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	party, err := ch.backend.GetPartyFromName(r.Context(), pname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired party", "name", name, "surname", surname, "party", pname, "user", user.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	village, err := ch.backend.GetVillageFromName(r.Context(), vname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired village", "name", name, "surname", surname, "party", party.Name, "user", user.Name, "village", vname, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	assoc := character.Associations{ CharacterID: char.ID, User: user, Party: party, Village: village }
	err = ch.backend.UpdateAssociations(r.Context(), assoc)
	if err != nil {
		log.Error("Failed to update character", "error", err, "name", name, "surname", surname, "user", uname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s", char.Name, char.Surname), http.StatusSeeOther)
}

func (ch *CharHandler) getAssociations(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	associations, err := ch.backend.GetAssociationsFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Failed to fetch character associations", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := charTpl.ExecuteTemplate(w, "associations", associations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) editAssociations(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	associations, err := ch.backend.GetAssociationsFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Failed to fetch character associations", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	model := character.AssociationsEdit{Name: name, Surname: surname, Associations: associations}
	if err := charTpl.ExecuteTemplate(w, "associations_edit", model); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) putAssociations(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	associations, err := ch.backend.GetAssociationsFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Could not find character", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	uname := r.FormValue("user")
	pname := r.FormValue("party")
	vname := r.FormValue("village")
	user, err := ch.backend.GetUserFromName(r.Context(), uname)
	if err != nil {
		log.Error("Could not find user when updating character", "error", err, "name", name, "surname", surname, "user", uname)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	party, err := ch.backend.GetPartyFromName(r.Context(), pname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired party", "name", name, "surname", surname, "party", pname, "user", user.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	village, err := ch.backend.GetVillageFromName(r.Context(), vname)
	if err != nil {
		log.Error("Failed to create new character: lookup failed for desired village", "name", name, "surname", surname, "party", party.Name, "user", user.Name, "village", vname, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	associations = character.Associations{ CharacterID: associations.CharacterID, User: user, Party: party, Village: village }
	err = ch.backend.UpdateAssociations(r.Context(), associations)
	if err != nil {
		log.Error("Failed to update character associations", "error", err, "name", name, "surname", surname, "user", uname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/associations", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) getIdentity(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	identity, err := ch.backend.GetIdentityFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Failed to fetch character identity", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := charTpl.ExecuteTemplate(w, "identity", identity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) editIdentity(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	identity, err := ch.backend.GetIdentityFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Failed to fetch character identity", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	model := character.IdentityEdit{Identity: identity}
	if err := charTpl.ExecuteTemplate(w, "identity_edit", model); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) putIdentity(w http.ResponseWriter, r *http.Request) {
	oldName := chi.URLParam(r, "name")
	oldSurname := chi.URLParam(r, "surname")
	identity, err := ch.backend.GetIdentityFromName(r.Context(), oldName, oldSurname)
	if err != nil {
		log.Error("Could not find character", "error", err, "name", oldName, "surname", oldSurname)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	ageStr := r.FormValue("age")
	age, err := strconv.ParseInt(ageStr, 10, 64)
	if err != nil {
		log.Error("Could not parse age of character as int when updating identity", "error", err, "name", oldName, "surname", oldSurname, "age", ageStr)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bname := r.FormValue("background")
	background, err := ch.backend.GetBackgroundFromTitle(r.Context(), bname)
	if err != nil {
		log.Error("Could not find background when updating character identity", "error", err, "name", oldName, "surname", oldSurname, "background", bname)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	gender := r.FormValue("gender")
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	portrait := r.FormValue("portrait")

	identity = character.Identity{
		CharacterID: identity.CharacterID,
		Gender: gender,
		Name: name,
		Surname: surname,
		Age: age,
		Portrait: portrait,
		Background: background,
	}

	err = ch.backend.UpdateIdentity(r.Context(), identity)
	if err != nil {
		log.Error("Failed to update character identity", "error", err, "name", name, "surname", surname, "old name", oldName, "old surname", oldSurname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/identity", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) getStatus(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	status, err := ch.backend.GetStatusFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Failed to fetch character status", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := charTpl.ExecuteTemplate(w, "status", status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) editStatus(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	status, err := ch.backend.GetStatusFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Failed to fetch character status", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	model := character.StatusEdit{Name: name, Surname: surname, Status: status}
	if err := charTpl.ExecuteTemplate(w, "status_edit", model); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) putStatus(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	status, err := ch.backend.GetStatusFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Could not find character", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	hpStr := r.FormValue("hp")
	hp, err := strconv.ParseInt(hpStr, 10, 64)
	strStr := r.FormValue("str")
	str, err := strconv.ParseInt(strStr, 10, 64)
	dexStr := r.FormValue("dex")
	dex, err := strconv.ParseInt(dexStr, 10, 64)
	willStr := r.FormValue("will")
	will, err := strconv.ParseInt(willStr, 10, 64)

	status.Hp = hp
	status.Str = str
	status.Dex = dex
	status.Will = will
	err = ch.backend.UpdateStatus(r.Context(), status)
	if err != nil {
		log.Error("Failed to update character status", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/status", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) putMaxStatus(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	status, err := ch.backend.GetStatusFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Could not find character", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	hpStr := r.FormValue("max_hp")
	hp, err := strconv.ParseInt(hpStr, 10, 64)
	strStr := r.FormValue("max_str")
	str, err := strconv.ParseInt(strStr, 10, 64)
	dexStr := r.FormValue("max_dex")
	dex, err := strconv.ParseInt(dexStr, 10, 64)
	willStr := r.FormValue("max_will")
	will, err := strconv.ParseInt(willStr, 10, 64)

	status.MaxHp = hp
	status.MaxStr = str
	status.MaxDex = dex
	status.MaxWill = will
	err = ch.backend.UpdateMaxStatus(r.Context(), status)
	if err != nil {
		log.Error("Failed to update character status", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/status", name, surname), http.StatusSeeOther)
}

func (ch *CharHandler) getTraits(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	traits, err := ch.backend.GetTraitsFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Failed to fetch character traits", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := charTpl.ExecuteTemplate(w, "traits", traits); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) editTraits(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")

	traits, err := ch.backend.GetTraitsFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Failed to fetch character traits", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	model := character.TraitsEdit{Name: name, Surname: surname, Traits: traits}
	if err := charTpl.ExecuteTemplate(w, "traits_edit", model); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ch *CharHandler) putTraits(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	traits, err := ch.backend.GetTraitsFromName(r.Context(), name, surname)
	if err != nil {
		log.Error("Could not find character", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

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

	traits = character.Traits{
		CharacterID: traits.CharacterID,
		Physique: physique,
		Skin: skin,
		Hair: hair,
		Face: face,
		Speech: speech,
		Clothing: clothing,
		Virtue: virtue,
		Vice: vice,
		Reputation: reputation,
		Misfortune: misfortune,
	}
	err = ch.backend.UpdateTraits(r.Context(), traits)
	if err != nil {
		log.Error("Failed to update character traits", "error", err, "name", name, "surname", surname)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/traits", name, surname), http.StatusSeeOther)
}
