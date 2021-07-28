package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"twitter/db"
	"twitter/jwt"
	"twitter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User password inVaild"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required ", 400)
		return
	}
	document, exist := db.IntentoLogin(t.Email, t.Password)
	if exist == false {
		http.Error(w, "User password inVaild exist=false", 400)
		return
	}
	jwtKey, err := jwt.GenerJWT(document)
	if err != nil {
		http.Error(w, "Occured  error intent to generate Token "+err.Error(), 400)
	}
	resp := models.ResquestLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
