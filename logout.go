package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		writeBad(w)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	sessionID := cookie.Value
	creden := strings.Split(sessionID, ":")
	user := creden[0]
	val := userList[user]

	sessionid := strconv.Itoa(rand.Intn(1000000007))
	hasher := sha1.New()
	hasher.Write([]byte(sessionid))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	cookieValue := val.Username + ":" + sha
	expire := time.Now().AddDate(0, 0, -1)
	newCookie := http.Cookie{Name: "SessionID", Value: cookieValue, Expires: expire, HttpOnly: true}
	http.SetCookie(w, &newCookie)

	userList[user] = User{val.Username, val.Password, val.Name, sha}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Success: 1, Message: "Logged out successfully!"})
}
