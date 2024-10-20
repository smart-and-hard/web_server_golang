package apiserver

import (
	"io"
	"net/http"
	"url-shortener/internal/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}
	s.configerRouter()

	if err := s.configerStore(); err != nil {
		return err
	}

	s.logger.Info("start api server")

	return http.ListenAndServe(s.config.BinAddr, s.router)
}

func (s *APIserver) configLogger() error {
	level, err := logrus.ParseLevel(s.config.logLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *APIserver) configerRouter() {
	s.router.HandleFunc("/hello", s.handlHello())
}

func (s *APIserver) configerStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil

}

func (s *APIserver) handlHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello oleg!")
	}
}
