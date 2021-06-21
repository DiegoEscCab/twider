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
	/* Esta sección es para el registro,login,ver y modificar del perfil.*/
	router.HandleFunc("/registro", middleW.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleW.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleW.ChequeoBD(middleW.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middleW.ChequeoBD(middleW.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	/*Esta sección es para publicar y eliminar un tweet */
	router.HandleFunc("/tweet", middleW.ChequeoBD(middleW.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middleW.ChequeoBD(middleW.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middleW.ChequeoBD(middleW.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
	/* Esta sección es para subir y obtener el avatar y banner del usuario */
	router.HandleFunc("/subirAvatar", middleW.ChequeoBD(middleW.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middleW.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middleW.ChequeoBD(middleW.ValidoJWT(routers.SubirBanners))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middleW.ChequeoBD(routers.ObtenerBanner)).Methods("GET")
	/*Esta sección es para poder seguir y dejar de seguir un perfil en la aplicación*/
	router.HandleFunc("/altaRelacion", middleW.ChequeoBD(routers.AltaRelacion)).Methods("POST")
	router.HandleFunc("/bajaRelacion", middleW.ChequeoBD(routers.BajaRelacion)).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middleW.ChequeoBD(routers.ConsultaRelacion)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
