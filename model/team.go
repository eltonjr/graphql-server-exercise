package model

import (
	"time"

	"github.com/graphql-go/graphql"
)

// Team represents a F1 scuderia
type Team struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
}

var TeamType *graphql.Object

func init() {
	TeamType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Team",
		Description: "A Formula 1 team",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Name:        "id",
				Type:        graphql.String,
				Description: "team's id",
			},
			"name": &graphql.Field{
				Name:        "name",
				Type:        graphql.String,
				Description: "team's name",
			},
			"createdAt": &graphql.Field{
				Name:        "createdAt",
				Type:        graphql.DateTime,
				Description: "teams creation date",
			},
		},
	})
}
