package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"cc/application"
	"cc/domain/entity"
)

type AppHandler struct {
	appInteractor application.AppInteractor
}

func NewAppHandler(appInteractor application.AppInteractor) *AppHandler {
	return &AppHandler{appInteractor}
}

func (handler *AppHandler) Add(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "application/json")
	var app entity.App
	err := json.NewDecoder(req.Body).Decode(&app)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(ErrResponse{Message: "Invalid Payload"})
		return
	}
	err = handler.appInteractor.CreateApp(app)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(ErrResponse{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("success"))
}

func (handler *AppHandler) FindAll(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	results, err := handler.appInteractor.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(ErrResponse{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(results)
}

func (handler *AppHandler) FindOne(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)

	sid := vars["id"]
	id, err := strconv.ParseUint(sid, 10, 64)
	encoder := json.NewEncoder(resp)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrResponse{err.Error()})
		return
	}
	result, err := handler.appInteractor.FindOne(id)
	if err != nil {
		encoder.Encode(ErrResponse{err.Error()})
		return
	}
	encoder.Encode(result)
}

func (handler *AppHandler) Delete(resp http.ResponseWriter, req *http.Request)  {
	encoder := json.NewEncoder(resp)
	vars := mux.Vars(req)
	sid := vars["id"]
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrResponse{err.Error()})
		return
	}
	err = handler.appInteractor.Delete(id)
	if err != nil {
		encoder.Encode(ErrResponse{err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("success"))
}

func (handler *AppHandler) Update(resp http.ResponseWriter, req *http.Request)  {
	var app entity.App
	encoder := json.NewEncoder(resp)
	resp.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&app)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrResponse{err.Error()})
		return
	}

	vars := mux.Vars(req)
	sid := vars["id"]
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrResponse{err.Error()})
		return
	}
	app.ID = id
	err = handler.appInteractor.Update(app)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrResponse{err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("success"))
}
