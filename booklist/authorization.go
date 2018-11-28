package booklist

import (
	"encoding/json"
	"net/http"
	"strings"
)

func isAuthorized(r *http.Request) bool {

	cookie, err := r.Cookie("SessionID")
	if err != nil {
		return false
	}

	mu.Lock()
	defer mu.Unlock()

	sessionID := cookie.Value
	creden := strings.Split(sessionID, ":")
	user := creden[0]
	sessionID = creden[1]

	expectedSessionID := userList[user].LastSessionID

	if expectedSessionID == sessionID {
		return true
	}
	return false
}

func authZ(h http.HandlerFunc, need bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAuthorized(r) != need {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Response{Success: 0, Message: "Not Authorized"})
			return
		}
		h.ServeHTTP(w, r)
	})
}
