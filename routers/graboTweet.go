package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/snetwork/bd"
	"github.com/snetwork/models"
)

/*GraboTweet graba el tweet en la bd*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
