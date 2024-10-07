package serve

import (
	"log"
	"net/http"
)

type Server struct {
	Addr string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server"))
}

func StartServer() {
	s := &Server{Addr: ":8080"}
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(s.Addr, s); err != nil {
		log.Fatal(err)
	}
}
