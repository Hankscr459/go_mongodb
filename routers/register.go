package routers

import (
	"encoding/json"
	"net/http"
	"twitter/db"
	"twitter/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error lost Data "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email required"+err.Error(), 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password at least 6 charactors"+err.Error(), 400)
		return
	}
	_, create, _ := db.CheckIsExistUser(t.Email)
	if create == true {
		http.Error(w, "This Email has been token", 400)
		return
	}

	_, status, err := db.InsertoRegister(t)
	if err != nil {
		http.Error(w, "Occurred error register user"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insert el register del user", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
