package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/graphql-server-exercise/db"
	"github.com/eltonjr/graphql-server-exercise/model"
)

type Router struct {
	driverDao     *db.DriverDao
	graphqlSchema graphql.Schema
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

func (router *Router) RegisterGraphQL(r *httprouter.Router) error {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getDriver": &graphql.Field{
				Type:        model.DriverType,
				Description: "Get driver by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: router.GetSingleDriverResolver,
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutations",
		Fields: graphql.Fields{
			"dummy": nil,
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		return fmt.Errorf("unable to create graphql schema: %v", err)
	}

	r.Handle(http.MethodGet, "/v2/drivers", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})

		json.NewEncoder(w).Encode(result)
	})

	return nil
}

func (router *Router) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Server is up!\n")
}
