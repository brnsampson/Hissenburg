package mux

import (
	"time"
	"bytes"
	"net/http"
	"html/template"
	"github.com/brnsampson/Hissenburg/models/view"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/charmbracelet/log"
)

const postLoginTempl = `
<!DOCTYPE html>
<html>

<head>
  <meta http-equiv="refresh" content="0; url='{{.}}'">
</head>

<body></body>

</html>`

func New(storePath string) (*chi.Mux, error) {
	rh, err := NewRootHandler()
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	router.Mount("/", rh)
	return router, nil
}


func SetCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string, path string, timeout time.Duration, revoke ...bool) {
	var maxAge int
	if len(revoke) > 0 && revoke[0] == true {
		maxAge = -1
	} else {
		maxAge = int(timeout.Seconds())
	}

	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		Secure:   r.TLS != nil,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	if path != "" {
		c.Path = path
	}
	http.SetCookie(w, c)
}

func LoginRedirect(w http.ResponseWriter, r *http.Request) {
	// Set cookie for previous request location
	SetCallbackCookie(w, r, "postLogin", r.URL.Path, "/", time.Hour)
	// Redirect to the login page
	http.Redirect(w, r, r.URL.Host+"/unauthorized", http.StatusFound)
}

func PostLoginRedirect(w http.ResponseWriter, r *http.Request) {
	// After users authenticate there is often a problem of cookies not
	// being set correstly if the server sends a 302 redirect. Serving a
	// 200 status and a short html doc which has the client redirect is
	// a reasonable solution.
	var redir string
	post, err := r.Cookie("postLogin")
	if err != nil {
		redir = r.Host
		log.Debug("No post login cookie found. Redirecting to the index instead", "redirect path", redir)
	} else {
		redir = r.Host + post.Value
		log.Debug("Post login cookie found.", "redirect path", redir)
		SetCallbackCookie(w, r, "postLogin", "", "/authenticate", 0, true)
	}

	// TODO: move this out of this function so we don't have to parse the template every call
	tmpl, err := template.New("postLogin").Parse(postLoginTempl)
	if err != nil {
		log.Error("Failed to parse post login redirect template", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	b := bytes.Buffer{}
	if err := tmpl.Execute(&b, redir); err != nil {
		log.Error("Error executing the post login redirect template", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//err = tmpl.Execute(w, redir)
	//if err != nil {
	//	log.Error("Failed to render post login redirect template", "error", err)
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	log.Debug("Rendered post login redirect template", "content", b.String())
	_, err = w.Write(b.Bytes())
	if err != nil {
		log.Error("Error writing post redirect template buffer to response buffer", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReturnPageViewUserRequired(w http.ResponseWriter, r *http.Request, b *bytes.Buffer) {
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

	ReturnPageView(w, r, b)
}

func ReturnPageView(w http.ResponseWriter, r *http.Request, b *bytes.Buffer) {
	user := ""
	uc, err := r.Cookie("username")
	if err == nil && uc.Valid() == nil {
		user = uc.Value
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	v := view.PageView{ User: user, Contents: template.HTML(b.String()) }
	if err := villageTpl.ExecuteTemplate(w, "page", v); err != nil {
		log.Error("Failed to template page view", "host", r.Host, "path", r.URL.Path, "view", v)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//type CharHandler struct {
//	chi.Router
//	backend data.CharBackend
//	itemBackend data.ItemBackend
//}
//
//func NewCharHandler(storePath string) (*CharHandler, error) {
//	backend, err := data.NewCharData(storePath)
//	if err != nil {
//		return nil, err
//	}
//
//	itemBackend, err := data.NewItemData()
//	if err != nil {
//		return nil, err
//	}
//
//	router := chi.NewRouter()
//	ch := CharHandler{router, backend, itemBackend}
//
//	router.Get("/", ch.listChars)
//	router.Post("/generate", ch.postGenerate)
//	router.Route("/{name}/{surname}", func(router chi.Router) {
//		router.Get("/", ch.getChar)
//		router.Post("/", ch.postChar)
//		router.Delete("/", ch.deleteChar)
//		router.Get("/identity", ch.getIdentity)
//		router.Get("/identity/edit", ch.editIdentity)
//		router.Put("/identity", ch.putIdentity)
//		router.Get("/status", ch.getStatus)
//		router.Get("/status/edit", ch.editStatus)
//		router.Put("/status", ch.putStatus)
//		router.Get("/traits", ch.getTraits)
//		router.Get("/traits/edit", ch.editTraits)
//		router.Put("/traits", ch.putTraits)
//		router.Get("/inventory", ch.getInventory)
//		router.Put("/inventory", ch.putInventory)
//		router.Route("/inventory/backpack", func(router chi.Router) {
//			router.Put("/", ch.putBackpack)
//			router.Get("/edit", ch.editBackpack)
//			router.Put("/{slot}", ch.putBackpackSlot)
//			router.Delete("/{slot}", ch.deleteBackpack)
//		})
//		router.Delete("/inventory/ground/{slot}", ch.deleteGround)
//	})
//
//	return &ch, nil
//}
//
//
//func (ch *CharHandler) listChars(w http.ResponseWriter, r *http.Request) {
//	chars, err := ch.backend.ListCharacters()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	if err := charTpl.ExecuteTemplate(w, "character_list", chars); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ch *CharHandler) postGenerate(w http.ResponseWriter, r *http.Request) {
//	char := character.New()
//	err := roll.RollChar(ch.backend, ch.itemBackend, &char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	_, err = ch.backend.CreateCharacter(&char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s", char.Name, char.Surname), http.StatusSeeOther)
//}
//
//func (ch *CharHandler) getChar(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	if err := charTpl.ExecuteTemplate(w, "character", char); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ch *CharHandler) postChar(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char := character.New()
//	char.Name = name
//	char.Surname = surname
//	_, err := ch.backend.CreateCharacter(&char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s", name, surname), http.StatusSeeOther)
//}
//
//func (ch *CharHandler) deleteChar(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	_, err := ch.backend.DeleteCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//    http.Redirect(w, r, "/character", http.StatusSeeOther)
//}
//
//func (ch *CharHandler) getIdentity(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	if err := charTpl.ExecuteTemplate(w, "identity", char); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ch *CharHandler) editIdentity(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	if err := charTpl.ExecuteTemplate(w, "edit_identity", char); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ch *CharHandler) putIdentity(w http.ResponseWriter, r *http.Request) {
//	oldName := chi.URLParam(r, "name")
//	oldSurname := chi.URLParam(r, "surname")
//
//	name := r.FormValue("name")
//	surname := r.FormValue("surname")
//	gender := r.FormValue("gender")
//	background := r.FormValue("background")
//	ageStr := r.FormValue("age")
//	age64, err := strconv.ParseUint(ageStr, 10, 16)
//	if err != nil {
//		http.Error(w, "Age must be an integer between 0 and 65535", http.StatusInternalServerError)
//	}
//	age := uint16(age64)
//
//	char, err := ch.backend.GetCharacter(oldName, oldSurname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	char.Name = name
//	char.Surname = surname
//	char.Gender = character.GenderFromString(gender)
//	char.Background = background
//	char.Age = age
//	char, err = ch.backend.UpdateCharacter(oldName, oldSurname, char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/identity", name, surname), http.StatusSeeOther)
//}
//
//func (ch *CharHandler) getInventory(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	if err := charTpl.ExecuteTemplate(w, "inventory", char); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ch *CharHandler) putInventory(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//
//	iStrings := r.Form["inventory"]
//	iKindStrings := r.Form["inventoryTypes"]
//	if len(iStrings) != len(iKindStrings) {
//		http.Error(w, "The list of items and list of item types were of different lengths", http.StatusInternalServerError)
//	}
//
//	items := make([]item.Item, 0)
//	for i, slot := range iStrings {
//		kind := item.TypeFromString(iKindStrings[i])
//		invItem, err := ch.itemBackend.GetItem(kind, slot)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//		}
//
//		items = append(items, invItem)
//	}
//
//	inv, err := inventory.Unpack(items)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	char.Inventory = inv
//
//	char, err = ch.backend.UpdateCharacter(name, surname, char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
//}
//
//func (ch *CharHandler) putBackpack(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//
//	iStrings := r.Form["item_name"]
//	iKindStrings := r.Form["item_type"]
//	if len(iStrings) != len(iKindStrings) {
//		http.Error(w, "The list of items and list of item types were of different lengths", http.StatusInternalServerError)
//	}
//
//	items := make([]item.Item, 0)
//	for i, slot := range iStrings {
//		kind := item.TypeFromString(iKindStrings[i])
//		invItem, err := ch.itemBackend.GetItem(kind, slot)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//		}
//
//		items = append(items, invItem)
//	}
//
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	for i, it := range items {
//		_, err := char.Inventory.SetBackpack(i, it)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//		}
//	}
//
//	char, err = ch.backend.UpdateCharacter(name, surname, char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
//}
//
//func (ch *CharHandler) putBackpackSlot(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	slot, err := strconv.Atoi(chi.URLParam(r, "slot"))
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	iName := r.Form["item_name"][slot]
//	iType := r.Form["item_type"][slot]
//
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	kind := item.TypeFromString(iType)
//	i, err := ch.itemBackend.GetItem(kind, iName)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	_, err = char.Inventory.SetBackpack(slot, i)
//
//	_, err = ch.backend.UpdateCharacter(name, surname, char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
//}
//
//func (ch *CharHandler) editBackpack(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	tl := item.ListItemTypes()
//	vm := char.Inventory.IntoView().IntoEditView(char.Name, char.Surname, tl)
//
//	if err := charTpl.ExecuteTemplate(w, "edit_backpack", vm); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//
//func (ch *CharHandler) deleteBackpack(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	slot, err := strconv.Atoi(chi.URLParam(r, "slot"))
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	old, err := char.Inventory.SetBackpack(slot, item.Empty())
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	char.Inventory.AddToGround(old)
//
//	char, err = ch.backend.UpdateCharacter(name, surname, char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
//}
//
//func (ch *CharHandler) deleteGround(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	slot, err := strconv.Atoi(chi.URLParam(r, "slot"))
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	if slot < 0 || slot >= len(char.Inventory.Ground) {
//		http.Error(w, "Addempted to delete a ground item out of range", http.StatusInternalServerError)
//	}
//	char.Inventory.Ground = append(char.Inventory.Ground[:slot], char.Inventory.Ground[slot + 1:]...)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	char, err = ch.backend.UpdateCharacter(name, surname, char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/inventory", name, surname), http.StatusSeeOther)
//}
//
//func (ch *CharHandler) getStatus(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	if err := charTpl.ExecuteTemplate(w, "status", char); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ch *CharHandler) editStatus(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	if err := charTpl.ExecuteTemplate(w, "edit_status", char); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ch *CharHandler) putStatus(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//
//	hpStr := r.FormValue("hp")
//	hp64, err := strconv.ParseUint(hpStr, 10, 8)
//	if err != nil {
//		http.Error(w, "HP must be an integer between 0 and 256 and less than MaxStr", http.StatusInternalServerError)
//	}
//	hp := uint8(hp64)
//
//	strStr := r.FormValue("str")
//	str64, err := strconv.ParseUint(strStr, 10, 8)
//	if err != nil {
//		http.Error(w, "Str must be an integer between 0 and 256 and less than MaxStr", http.StatusInternalServerError)
//	}
//	str := uint8(str64)
//
//	dexStr := r.FormValue("dex")
//	dex64, err := strconv.ParseUint(dexStr, 10, 8)
//	if err != nil {
//		http.Error(w, "Dex must be an integer between 0 and 256 and less than MaxDex", http.StatusInternalServerError)
//	}
//	dex := uint8(dex64)
//
//	willStr := r.FormValue("will")
//	will64, err := strconv.ParseUint(willStr, 10, 8)
//	if err != nil {
//		http.Error(w, "Will must be an integer between 0 and 256 and less than MaxWill", http.StatusInternalServerError)
//	}
//	will := uint8(will64)
//
//	maxhpStr := r.FormValue("maxhp")
//	maxhp64, err := strconv.ParseUint(maxhpStr, 10, 8)
//	if err != nil {
//		http.Error(w, "MaxHP must be an integer between 0 and 256", http.StatusInternalServerError)
//	}
//	maxhp := uint8(maxhp64)
//
//	maxstrStr := r.FormValue("maxstr")
//	maxstr64, err := strconv.ParseUint(maxstrStr, 10, 8)
//	if err != nil {
//		http.Error(w, "Str must be an integer between 0 and 256", http.StatusInternalServerError)
//	}
//	maxstr := uint8(maxstr64)
//
//	maxdexStr := r.FormValue("maxdex")
//	maxdex64, err := strconv.ParseUint(maxdexStr, 10, 8)
//	if err != nil {
//		http.Error(w, "Dex must be an integer between 0 and 256", http.StatusInternalServerError)
//	}
//	maxdex := uint8(maxdex64)
//
//	maxwillStr := r.FormValue("maxwill")
//	maxwill64, err := strconv.ParseUint(maxwillStr, 10, 8)
//	if err != nil {
//		http.Error(w, "Will must be an integer between 0 and 256", http.StatusInternalServerError)
//	}
//	maxwill := uint8(maxwill64)
//
//	status := status.New()
//	status.HP = hp
//	status.Str = str
//	status.Dex = dex
//	status.Will = will
//	status.MaxHP = maxhp
//	status.MaxStr = maxstr
//	status.MaxDex = maxdex
//	status.MaxWill = maxwill
//
//	err = ch.backend.UpdateStatus(name, surname, status)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/status", name, surname), http.StatusSeeOther)
//}
//
//func (ch *CharHandler) getTraits(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	if err := charTpl.ExecuteTemplate(w, "traits", char); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ch *CharHandler) editTraits(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	if err := charTpl.ExecuteTemplate(w, "edit_traits", char); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ch *CharHandler) putTraits(w http.ResponseWriter, r *http.Request) {
//	name := chi.URLParam(r, "name")
//	surname := chi.URLParam(r, "surname")
//
//	physique := r.FormValue("physique")
//	skin := r.FormValue("skin")
//	hair := r.FormValue("hair")
//	face := r.FormValue("face")
//	speech := r.FormValue("speech")
//	clothing := r.FormValue("clothing")
//	virtue := r.FormValue("virtue")
//	vice := r.FormValue("vice")
//	reputation := r.FormValue("reputation")
//	misfortune := r.FormValue("misfortune")
//
//
//	traits := character.NewTraits()
//	traits.Physique = physique
//	traits.Skin = skin
//	traits.Hair = hair
//	traits.Face = face
//	traits.Speech = speech
//	traits.Clothing = clothing
//	traits.Virtue = virtue
//	traits.Vice = vice
//	traits.Reputation = reputation
//	traits.Misfortune = misfortune
//
//	char, err := ch.backend.GetCharacter(name, surname)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	char.Traits = traits
//
//	char, err = ch.backend.UpdateCharacter(name, surname, char)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//    http.Redirect(w, r, fmt.Sprintf("/character/%s/%s/traits", name, surname), http.StatusSeeOther)
//}
