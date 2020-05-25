package routers

import (
	"encoding/json"
	"net/http"

	"github.com/snetwork/bd"

	"github.com/snetwork/models"
)

/*Registro de usuarios*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return

	}
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido por favor", 400)
	}
	if len(t.Password) < 0 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido por favor", 400)
		return
	}
	//Llama si ya existe el email en la bd
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al registrar un usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado el registro del usuario"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
