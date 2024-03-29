package server

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maxheckel/scare-me-to-sleep/internal/config"
	"github.com/maxheckel/scare-me-to-sleep/internal/db"
	"github.com/rs/cors"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	db      *gorm.DB
	config  *config.Config
	prompts db.PromptsRepository
}

func NewApp(cfg *config.Config) (*App, error) {
	database, err := db.Connect(cfg.DBFile)
	if err != nil {
		return nil, err
	}
	return &App{
		db:      database,
		config:  cfg,
		prompts: db.NewPromptsRepository(database),
	}, nil
}

func (a App) Start() {
	rtr := mux.NewRouter()

	rtr.Handle("/healthcheck", RecoverWrap(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})))

	rtr.Handle("/api/day", RecoverWrap(http.HandlerFunc(a.GetDay)))

	rtr.PathPrefix("/").Handler(RecoverWrap(spaHandler{
		staticPath: "frontend/dist",
		indexPath:  "index.html",
	}))

	http.Handle("/", rtr)
	log.Println("Listening...")

	if err := http.ListenAndServe(":80", cors.Default().Handler(rtr)); err != nil {
		fmt.Println(err)
	}
}

func RecoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
