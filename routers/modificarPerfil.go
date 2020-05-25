package routers

import (
	"encoding/json"
	"net/http"

	"github.com/snetwork/bd"
	"github.com/snetwork/models"
)

/*ModificarPerfil llama a bd al metodo modificar registro y hace la verificacion de errores	*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error al intentar modificar"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
	}
	w.WriteHeader(http.StatusCreated)
}
