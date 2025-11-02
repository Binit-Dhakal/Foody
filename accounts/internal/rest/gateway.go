package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Binit-Dhakal/Foody/accounts/internal/application"
	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/go-chi/chi/v5"
)

type AccountHandler struct {
	mux     *chi.Mux
	userSvc application.UserService
	authSvc application.AuthService
}

func NewAccountHandler(mux *chi.Mux, userSvc application.UserService, authSvc application.AuthService) AccountHandler {
	return AccountHandler{
		mux:     mux,
		userSvc: userSvc,
		authSvc: authSvc,
	}
}

func (h *AccountHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	req.Validate()
	if req.Validator.HasErrors() {
		http.Error(w, fmt.Sprintf("Validation error: %+v", req.Validator), http.StatusUnprocessableEntity)
		return
	}

	err := h.userSvc.RegisterCustomer(r.Context(), &req)
	if err != nil {
		http.Error(w, "failed to register user:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AccountHandler) RegisterResturant(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterResturantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	req.Validate()
	if req.Validator.HasErrors() {
		http.Error(w, fmt.Sprintf("Validation error: %+v", req.Validator), http.StatusUnprocessableEntity)
		return
	}

	err := h.userSvc.RegisterVendor(r.Context(), &req)
	if err != nil {
		http.Error(w, "failed to register resturant:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AccountHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	req.Validate()
	if req.Validator.HasErrors() {
		http.Error(w, fmt.Sprintf("Validation error: %+v", req.Validator), http.StatusUnprocessableEntity)
		return
	}

	_, err := h.authSvc.Login(r.Context(), &req)
	if err != nil {
		http.Error(w, "failed to authenticate:"+err.Error(), http.StatusInternalServerError)
		return
	}
}
