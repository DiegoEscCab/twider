package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DiegoEscCab/twider/bd"
	"github.com/DiegoEscCab/twider/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "NO se ha  logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
