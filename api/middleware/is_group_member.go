package middleware

import (
	"log"
	"net/http"
	"strconv"

	"git.01.alem.school/Kusbek/social-network/usecase/group"
)

func IsGroupMember(groupService group.UseCase, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		groupID, err := strconv.Atoi(r.URL.Query().Get("group_id"))
		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		isMember, err := groupService.IsGroupMember(r.Context().Value(UserID).(int), groupID)
		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if !isMember {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("You are not member of this group"))
			return
		}
		next(w, r)
	})

}
