package apiserver

import (
	"github.com/JadesHeart/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start server
func (s APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	} // create logger

	s.logger.Info("Server start...") // Test logger!

	//if err := s.configureStore(); err != nil {
	//	return err
	//}

	s.configureRouter() //

	return http.ListenAndServe("localhost"+s.config.BindAddr, s.router)

}

// Logger configuration
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

// Router configuration
func (s *APIServer) configureRouter() {
	conn := s.router
	conn.HandleFunc("/hello", s.handHello())
}

// Testing request
func (s *APIServer) handHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Привет, Мир!")
	}
}
