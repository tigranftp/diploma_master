package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"nn_auto/db"
)

type API struct {
	router *mux.Router
	store  db.Storage
}

func NewAPI(storage db.Storage) (*API, error) {
	api := new(API)
	router := mux.NewRouter()
	router.HandleFunc("/logging", api.loggingHandler())
	router.HandleFunc("/ping", api.pingHandler())
	router.HandleFunc("/register", api.registerHandler())
	router.HandleFunc("/create_model", api.createModelHandler())
	api.router = router
	api.store = storage
	return api, nil
}

func (a *API) Start() error {
	//a.configureRouter()
	//a.configureDB()
	//if err := a.store.Open(); err != nil {
	//	return err
	//}
	return http.ListenAndServe(":36363", a.router)
}

func (a *API) Stop() {
	fmt.Println("Stopping API...")
	a.store.Close()
	fmt.Println("API stopped...")
}
