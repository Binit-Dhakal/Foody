package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"runtime/debug"

	"github.com/Binit-Dhakal/Foody/accounts/internal/application"
	"github.com/Binit-Dhakal/Foody/accounts/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/cookies"
	ctxutil "github.com/Binit-Dhakal/Foody/internal/ctxutils"
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
	fmt.Println(ctxutil.GetContext(r.Context(), ctxutil.UserContextKey))
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

	token, err := h.authSvc.LoginUser(r.Context(), &req)
	if err != nil {
		debug.PrintStack()
		http.Error(w, "failed to authenticate:"+err.Error(), http.StatusInternalServerError)
		return
	}

	accessCookie := &http.Cookie{
		Name:     "accessToken",
		Value:    token.Token,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	cookies.Write(w, accessCookie)

	refreshCookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    token.RefreshToken,
		Expires:  time.Now().Add(15 * 24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	cookies.Write(w, refreshCookie)

	w.WriteHeader(http.StatusOK)
}

func (h *AccountHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	err = h.authSvc.LogoutUser(r.Context(), cookie.Value)
	if err != nil {
		http.Error(w, "failed to logout", http.StatusInternalServerError)
		return
	}

	clearCookie := func(name string) {
		http.SetCookie(w, &http.Cookie{
			Name:     name,
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		})
	}

	clearCookie("accessToken")
	clearCookie("refreshToken")

	w.WriteHeader(http.StatusOK)
}
