package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/Binit-Dhakal/Foody/cmd/foody/internal/domain"
	"github.com/Binit-Dhakal/Foody/internal/cookies"
	ctxutil "github.com/Binit-Dhakal/Foody/internal/ctxutils"
	"github.com/Binit-Dhakal/Foody/internal/jwtutil"
)

func AuthenticateMiddleware(accountsClient domain.AccountRepository, secretKey []byte) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken, _ := cookies.Read(r, "accessToken")

			claims, err := jwtutil.VerifyToken(accessToken, secretKey)
			if nil == err {
				ctx := ctxutil.AddContext(r.Context(),
					ctxutil.UserContextKey, claims.UserID,
					ctxutil.RoleContextKey, claims.RoleID,
				)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			if errors.Is(err, jwtutil.ErrTokenExpired) {
				refreshToken, _ := cookies.Read(r, "refreshToken")
				if refreshToken == "" {
					http.Error(w, "unauthorized", http.StatusUnauthorized)
					return
				}

				newToken, err := accountsClient.RefreshToken(r.Context(), refreshToken)
				if err != nil {
					http.Error(w, "unauthorized", http.StatusUnauthorized)
					return
				}

				accessCookie := &http.Cookie{
					Name:     "accessToken",
					Value:    newToken.AccessToken,
					Expires:  time.Now().Add(24 * time.Hour),
					Path:     "/",
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteLaxMode,
				}
				cookies.Write(w, accessCookie)

				refreshCookie := &http.Cookie{
					Name:     "refreshToken",
					Value:    newToken.RefreshToken,
					Expires:  time.Now().Add(15 * 24 * time.Hour),
					Path:     "/",
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteLaxMode,
				}
				cookies.Write(w, refreshCookie)

				claims, _ := jwtutil.VerifyToken(accessToken, secretKey)

				ctx := ctxutil.AddContext(r.Context(),
					ctxutil.UserContextKey, claims.UserID,
					ctxutil.RoleContextKey, claims.RoleID,
				)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			ctx := ctxutil.AddContext(r.Context(),
				ctxutil.UserContextKey, "",
				ctxutil.RoleContextKey, "",
			)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}
