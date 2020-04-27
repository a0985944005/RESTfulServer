package main

import (
	"log"
	"net/http"

	// _ "../PlatformServer/cache"
	_ "../PlatformServer/database"

	routes "../PlatformServer/router"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)
	log.Print("Server port:3000 listening")
}
