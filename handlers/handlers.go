package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/DiegoEscCab/twider/middleW"
	"github.com/DiegoEscCab/twider/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar el servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleW.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleW.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleW.ChequeoBD(middleW.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middleW.ChequeoBD(middleW.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middleW.ChequeoBD(middleW.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middleW.ChequeoBD(middleW.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
