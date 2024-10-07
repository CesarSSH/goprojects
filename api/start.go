package api

import (
	"log"
	"net/http"
)

func StartService() {
	api := &Api{Addr: ":8080"}

	//Initialize the ServeMux
	mux := http.NewServeMux()
	log.Println("Starting server on :8080")

	srv := &http.Server{
		Addr:    api.Addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.GetUsersHandler)
	mux.HandleFunc("POST /users", api.CreateUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
