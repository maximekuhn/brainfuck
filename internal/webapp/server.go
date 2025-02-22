package webapp

import "net/http"

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	fs := http.FileServer(http.Dir("internal/webapp/ui/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	logger := NewLogger()
	service := newService()

	handler := NewHandler(logger, service)
	http.Handle("/", handler)

	return http.ListenAndServe(":6006", nil)
}
