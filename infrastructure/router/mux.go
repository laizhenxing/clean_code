package router

import (
	"github.com/gorilla/mux"
)

func initRouter(muxR *mux.Router)  {
	ctrl := getAppHandler()

	muxR.HandleFunc("/v1/app/{id}", ctrl.FindOne).Methods("GET")
	muxR.HandleFunc("/v1/apps", ctrl.FindAll).Methods("GET")
	muxR.HandleFunc("/v1/app/{id}", ctrl.Update).Methods("POST")
	muxR.HandleFunc("/v1/app/{id}", ctrl.Delete).Methods("DELETE")
	muxR.HandleFunc("/v1/app", ctrl.Add).Methods("POST")
}

func OriginHandler() *mux.Router {
	httpSrv := mux.NewRouter()

	initRouter(httpSrv)

	return httpSrv
}