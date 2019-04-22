package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/graphql-server-exercise/model"
)

func (router *Router) GetDriversHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	drivers, err := router.driverDao.GetDrivers()
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
	driver, err := router.driverDao.GetSingleDriver(id)
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
	var driver *model.Driver
	err := json.NewDecoder(r.Body).Decode(driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error receiving driver: %v", err)
		return
	}
	defer r.Body.Close()

	driver, err = router.driverDao.CreateDriver(driver)
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

	driver.ID = p.ByName("id")

	driver, err = router.driverDao.UpdateDriver(driver)
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
	driver, err := router.driverDao.DeleteDriver(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting driver with id '%s': %v", id, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(driver)
}
