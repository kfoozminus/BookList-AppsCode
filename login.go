package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	user, pass, logok := r.BasicAuth()

	if logok == true {

		if val, ok := userList[user]; ok {

			if pass == val.Password {

				sessionid := strconv.Itoa(rand.Intn(1000000007))
				hasher := sha1.New()
				hasher.Write([]byte(sessionid))
				sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

				cookieValue := val.Username + ":" + sha
				expire := time.Now().AddDate(0, 0, 1)
				cookie := http.Cookie{Name: "SessionID", Value: cookieValue, Expires: expire, HttpOnly: true}
				http.SetCookie(w, &cookie)

				userList[user] = User{val.Username, val.Password, val.Name, sha}
				json.NewEncoder(w).Encode(Response{Success: 1, Message: "Login Successful"})

			} else {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(Response{Success: 0, Message: "Password doesn't match"})
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{Success: 0, Message: "User Not Found"})
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "Login Unsuccessgul"})
	}
}
