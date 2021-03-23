package middleware

import (
	"context"
	"log"
	"net/http"

	"git.01.alem.school/Kusbek/social-network/usecase/session"
)

type key int

const (
	UserID key = iota
)

func Auth(sessionService session.UseCase, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		u, err := sessionService.GetSession(cookie.Value)
		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		ctx := context.WithValue(r.Context(), UserID, u.ID)
		next(w, r.WithContext(ctx))
	})

}
