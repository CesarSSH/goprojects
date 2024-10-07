package serve

import (
	"log"
	"net/http"
)

type Server struct {
	Addr string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello from the server"))
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Index page"))
			return
		case "/users":
			w.Write([]byte("Users page"))
			return
		case "/admins":
			w.Write([]byte("Admins page"))
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Page Not Found"))
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Invalid method"))

	}
}

func StartServer() {
	s := &Server{Addr: ":8080"}
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(s.Addr, s); err != nil {
		log.Fatal(err)
	}
}
