package user

import (
	
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter,r *http.Request) {

    // do something
}

func (h *Handler) handleRegister(w http.ResponseWriter,r *http.Request) {

    // do something

	//get JSON payload
	//check if the user exists
	// if it doesnt we create the new user
}