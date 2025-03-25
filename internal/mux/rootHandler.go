package mux

import (
	"bytes"
	"database/sql"
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"time"
)

var assets = http.FileServer(http.Dir("./assets"))
var indexTpl = template.Must(template.ParseFiles("./templates/index.html"))
var loginTpl = template.Must(template.ParseFiles("./templates/user/login.html", "./templates/user/loginRedirect.html", "./templates/user/logout.html", "./templates/common.html"))
var villageTpl = template.Must(template.ParseFiles("./templates/village/village.html", "./templates/common.html"))
var partyTpl = template.Must(template.ParseFiles("./templates/party/party.html", "./templates/common.html"))

type RootHandler struct {
	chi.Router
	backend data.MetaRepo
}

func NewRootHandler() (*RootHandler, error) {
	chandler, err := NewCharHandler()
	if err != nil {
		return nil, err
	}
	invHandler, err := NewInvHandler()
	if err != nil {
		return nil, err
	}
	itemHandler, err := NewItemHandler()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", "file:./db/hiss.sqlite3?_foreign_keys=1")
	if err != nil {
		log.Error("Error while creating backend for RootHandler")
		return nil, err
	}

	backend := data.NewRepo(db)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	rh := RootHandler{router, backend}

	router.Get("/", rh.getIndex)
	router.Get("/login", rh.getLogin)
	router.Get("/loginRedirect", rh.getLoginRedirect)
	router.Post("/login", rh.postLogin)
	router.Post("/user", rh.postUser)
	router.Get("/logout", rh.getLogout)
	router.Get("/unauthorized", rh.getUnauthorized)
	router.Route("/village", func(router chi.Router) {
		router.Get("/", rh.listVillage)
		router.Post("/", rh.postVillage)
		router.Get("/{name}", rh.getVillage)
	})
	router.Route("/party", func(router chi.Router) {
		router.Get("/", rh.listParties)
		router.Post("/", rh.postParty)
		router.Get("/{name}", rh.getParty)
	})

	router.Mount("/character", chandler)
	router.Mount("/inventory", http.StripPrefix("/inventory", invHandler))
	router.Mount("/item", http.StripPrefix("/item", itemHandler))
	router.Mount("/assets", http.StripPrefix("/assets", assets))

	return &rh, nil
}

func (rh *RootHandler) getIndex(w http.ResponseWriter, r *http.Request) {

	b := bytes.Buffer{}
	if err := indexTpl.Execute(&b, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ReturnPageView(w, r, &b)
}

func (rh *RootHandler) getLogin(w http.ResponseWriter, r *http.Request) {
	if err := loginTpl.ExecuteTemplate(w, "login", nil); err != nil {
		log.Error("Error executing the login teamplate")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rh *RootHandler) getUnauthorized(w http.ResponseWriter, r *http.Request) {

	b := bytes.Buffer{}
	if err := loginTpl.ExecuteTemplate(&b, "login", nil); err != nil {
		log.Error("Error executing the login teamplate")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ReturnPageView(w, r, &b)
}

func (rh *RootHandler) getLoginRedirect(w http.ResponseWriter, r *http.Request) {
	// After users authenticate there is often a problem of cookies not
	// being set correstly if the server sends a 302 redirect. Serving a
	// 200 status and a short html doc which has the client redirect is
	// a reasonable solution.
	redir := "/"
	post, err := r.Cookie("postLogin")
	if err != nil {
		//redir = r.Host
		log.Debug("No post login cookie found. Redirecting to the index instead", "redirect path", redir)
	} else {
		redir = post.Value
		log.Debug("Post login cookie found", "redirect path", redir)
		SetCallbackCookie(w, r, "postLogin", "", "/", 0, true)
	}

	w.Header().Set("HX-Redirect", redir)
	//http.Redirect(w, r, redir, http.StatusSeeOther)
	if err := loginTpl.ExecuteTemplate(w, "loginRedirect", redir); err != nil {
		log.Error("Error executing the post login redirect teamplate")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rh *RootHandler) postLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username == "" {
		log.Error("Error creating new user: username was blank.")
		http.Error(w, "User name was blank", http.StatusInternalServerError)
		return
	}

	user, err := rh.backend.GetUserFromName(r.Context(), username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	SetCallbackCookie(w, r, "username", user.Name, "/", time.Second*86400)

	http.Redirect(w, r, "/loginRedirect", http.StatusSeeOther)
}

func (rh *RootHandler) getLogout(w http.ResponseWriter, r *http.Request) {
	SetCallbackCookie(w, r, "username", "", "/", 1, true)
	w.Header().Set("HX-Redirect", "/")
	if err := loginTpl.ExecuteTemplate(w, "logoutRedirect", ""); err != nil {
		log.Error("Error executing the post logout redirect template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rh *RootHandler) postUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username == "" {
		log.Error("Error creating new user: username was blank.")
		http.Error(w, "User name was blank", http.StatusInternalServerError)
		return
	}

	user, err := rh.backend.CreateUser(r.Context(), username)
	if err != nil {
		log.Error("Error creating new user", "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	SetCallbackCookie(w, r, "username", user.Name, "/", time.Second*86400)
	PostLoginRedirect(w, r)
}

func (rh *RootHandler) listVillage(w http.ResponseWriter, r *http.Request) {
	b := bytes.Buffer{}
	villages, err := rh.backend.ListVillages(r.Context())
	if err != nil {
		log.Error("Error listing villages from database!", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := villageTpl.ExecuteTemplate(&b, "village_browser", villages); err != nil {
		log.Error("Error executing the village overview teamplate")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ReturnPageView(w, r, &b)
}

func (rh *RootHandler) postVillage(w http.ResponseWriter, r *http.Request) {
	vname := r.FormValue("name")
	if vname == "" {
		log.Error("Error creating new village: name was blank.")
		http.Error(w, "Village name was blank", http.StatusInternalServerError)
		return
	}

	village, err := rh.backend.CreateVillage(r.Context(), vname)
	if err != nil {
		log.Error("Error creating new village", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/village/"+village.Name, http.StatusSeeOther)
}

func (rh *RootHandler) getVillage(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "new" {
		if err := villageTpl.ExecuteTemplate(w, "village_new", nil); err != nil {
			log.Error("Error executing the village creation teamplate")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		village, err := rh.backend.GetVillageFromName(r.Context(), name)
		if err != nil {
			log.Error("Could not fine village", "name", name)
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		if err := villageTpl.ExecuteTemplate(w, "village_view", village); err != nil {
			log.Error("Error executing the village teamplate")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (rh *RootHandler) listParties(w http.ResponseWriter, r *http.Request) {
	b := bytes.Buffer{}
	if err := partyTpl.ExecuteTemplate(&b, "party_browser", nil); err != nil {
		log.Error("Error executing the party overview teamplate")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ReturnPageView(w, r, &b)
}

func (rh *RootHandler) postParty(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("name")
	if name == "" {
		log.Error("Error creating new party: name was blank.")
		http.Error(w, "Party name was blank", http.StatusInternalServerError)
		return
	}

	toCreate := sqlc.Party{Name: name, Description: description}
	party, err := rh.backend.CreateParty(r.Context(), toCreate)
	if err != nil {
		log.Error("Error creating new party", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/party/"+party.Name, http.StatusSeeOther)
}

func (rh *RootHandler) getParty(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "new" {
		if err := partyTpl.ExecuteTemplate(w, "party_new", nil); err != nil {
			log.Error("Error executing the party creation teamplate")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		party, err := rh.backend.GetPartyFromName(r.Context(), name)
		if err != nil {
			log.Error("Could not find party", "name", name)
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		if err := partyTpl.ExecuteTemplate(w, "party_view", party); err != nil {
			log.Error("Error executing the party teamplate")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
