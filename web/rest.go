package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/graphql-server-exercise/db"
	"github.com/eltonjr/graphql-server-exercise/model"
)

func (router *Router) GetDriversHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	drivers, err := db.GetDrivers(0, 10) // TODO move this to somewhere else
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting drivers: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(drivers)
}

func (router *Router) GetSingleDriverHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	idHex, err := primitive.ObjectIDFromHex(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid driver id: %v", err)
		return
	}

	driver, err := db.GetSingleDriver(idHex)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting driver with id '%s': %v", id, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(driver)
}

func (router *Router) CreateDriverHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	driver := &model.Driver{}
	err := json.NewDecoder(r.Body).Decode(driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error receiving driver: %v", err)
		return
	}
	defer r.Body.Close()

	driver, err = db.CreateDriver(driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error saving driver: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(driver)
}

func (router *Router) UpdateDriverHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var driver *model.Driver
	err := json.NewDecoder(r.Body).Decode(driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error receiving driver: %v", err)
		return
	}
	defer r.Body.Close()

	id, err := primitive.ObjectIDFromHex(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid driver id: %v", err)
		return
	}
	driver.ID = id

	driver, err = db.UpdateDriver(driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error updating driver: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(driver)
}

func (router *Router) DeleteDriverHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	idHex, err := primitive.ObjectIDFromHex(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid driver id: %v", err)
		return
	}

	err = db.DeleteDriver(idHex)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting driver with id '%s': %v", id, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("")
}
