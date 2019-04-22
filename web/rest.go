package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (router *Router) GetDriversHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	driver, err := router.driverDao.GetDrivers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting drivers: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(driver)
}

func (router *Router) GetSingleDriverHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	json.NewEncoder(w).Encode("TODO: Method Not Implemented")
}

func (router *Router) CreateDriverHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	json.NewEncoder(w).Encode("TODO: Method Not Implemented")
}

func (router *Router) UpdateDriverHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	json.NewEncoder(w).Encode("TODO: Method Not Implemented")
}

func (router *Router) DeleteDriverHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	json.NewEncoder(w).Encode("TODO: Method Not Implemented")
}
