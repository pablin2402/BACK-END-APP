package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/snetwork/middlew"
	"github.com/snetwork/routers"
)

/*Manejadores Se ejecuta al llamar a la api, RUTAS*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweets))).Methods("DELETE")
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")
	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(routers.AltaRelacion)).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(routers.BajaRelacion)).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(routers.ConsultaRelacion)).Methods("GET")
	router.HandleFunc("/listaUsuario", middlew.ChequeoBD(routers.ListaUsuarios)).Methods("GET")
	router.HandleFunc("/leoTweetSeguidores", middlew.ChequeoBD(routers.LeoTweetsSeguidores)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"

	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
