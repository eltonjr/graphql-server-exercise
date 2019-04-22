package model

import "github.com/graphql-go/graphql"

// Driver represents a F1 driver
type Driver struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

var DriverType *graphql.Object

func init() {
	DriverType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Driver",
		Description: "A Formula 1 Driver",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Name:        "id",
				Type:        graphql.String,
				Description: "Driver's id",
			},
			"name": &graphql.Field{
				Name:        "name",
				Type:        graphql.String,
				Description: "Driver's name",
			},
			"country": &graphql.Field{
				Name:        "country",
				Type:        graphql.String,
				Description: "Driver's birth place",
			},
		},
	})
}
