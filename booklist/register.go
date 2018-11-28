package booklist

import (
	"encoding/json"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		writeBad(w)
		return
	}

	if user.Username == "" || user.Password == "" || user.Name == "" {
		writeBad(w)
		return
	}

	if _, ok := userList[user.Username]; ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "Username already exists"})
		return
	}

	userList[user.Username] = user
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Success: 1, Message: "Successfully registered"})
}
