package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/graphql-server-exercise/model"
)

type Router struct {
	graphqlSchema graphql.Schema
}

func NewRouter() *Router {
	return &Router{}
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
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: router.GetSingleDriverResolver,
			},
			"getDrivers": &graphql.Field{
				Type:        graphql.NewList(model.DriverType),
				Description: "Get list of drivers",
				Resolve:     router.GetDriversResolver,
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutations",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        model.DriverType,
				Description: "Create a driver",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: router.CreateDriverResolver,
			},
			"update": &graphql.Field{
				Type:        model.DriverType,
				Description: "Update driver by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: router.UpdateDriverResolver,
			},
			"delete": &graphql.Field{
				Type:        model.DriverType,
				Description: "Delete driver by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: router.DeleteDriverResolver,
			},
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
