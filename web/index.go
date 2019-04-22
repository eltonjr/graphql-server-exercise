package web

import (
	"fmt"
	"net/http"

	"github.com/eltonjr/graphql-server-exercise/db"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	driverDao *db.DriverDao
}

func NewRouter(driverDao *db.DriverDao) *Router {
	return &Router{
		driverDao: driverDao,
	}
}

func (router *Router) RegisterIndex(r *httprouter.Router) {
	r.GET("/", router.Index)
}

func (router *Router) RegisterREST(r *httprouter.Router) {
	r.GET("/v1/drivers", router.GetDriversHandler)
	r.GET("/v1/drivers/:id", router.GetSingleDriverHandler)
	r.POST("/v1/drivers", router.CreateDriverHandler)
	r.PUT("/v1/drivers/:id", router.UpdateDriverHandler)
	r.DELETE("/v1/drivers/:id", router.DeleteDriverHandler)
}

func (router *Router) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Server is up!\n")
}
