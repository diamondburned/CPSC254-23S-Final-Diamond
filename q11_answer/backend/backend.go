package backend

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"q11_answer/bank"
	"strings"
)

type BankService interface {
	bank.BankService
	bank.UserService
}

type handler struct{ s BankService }

func New(s BankService) http.Handler {
	h := handler{s: s}

	r := http.NewServeMux()
	r.HandleFunc("/login", h.login)
	r.HandleFunc("/me", h.me)
	r.HandleFunc("/", h.index)

	return r
}

func (h handler) login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "frontend/login.html")
	case http.MethodPost:
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := h.s.Login(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "basic", // mega secure :)
			Value: base64.StdEncoding.EncodeToString([]byte(username + ":" + password)),
			Path:  "/",
		})

		http.Redirect(w, r, "/", http.StatusFound)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h handler) authRequest(r *http.Request) (*bank.User, error) {
	cookie, err := r.Cookie("basic")
	if err != nil {
		return nil, err
	}

	cred, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, err
	}

	username, password, _ := strings.Cut(string(cred), ":")
	if err != nil {
		return nil, err
	}

	u, err := h.s.Login(username, password)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (h handler) index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "frontend/index.html")
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h handler) me(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u, err := h.authRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(u)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
