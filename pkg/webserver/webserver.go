package webserver

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type method string

var POST method = "POST"
var GET method = "GET"

type WebServer struct {
	Router        chi.Router
	Handlers      []Handler
	WebServerPort string
}

type Handler struct {
	Method  method
	Handler http.HandlerFunc
	Path    string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]Handler, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, m method, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, Handler{
		Method:  m,
		Handler: handler,
		Path:    path,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, wsh := range s.Handlers {
		switch wsh.Method {
		case POST:
			s.Router.Post(wsh.Path, wsh.Handler)
		case GET:
			s.Router.Get(wsh.Path, wsh.Handler)
		}
	}
	err := http.ListenAndServe(s.WebServerPort, s.Router)
	if err != nil {
		log.Fatal(err)
	}
}
