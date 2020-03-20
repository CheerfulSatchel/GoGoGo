package server

import (
	"net/http"
)

func StartServer() {
	serverRoutes := AllRoutes()
	router := NewRouter(serverRoutes)
	http.ListenAndServe(":8080", router)
}
