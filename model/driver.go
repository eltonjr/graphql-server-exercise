package model

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Driver represents a F1 driver
type Driver struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Country string             `json:"country" bson:"country"`
	Seasons []DriverSeason     `json:"seasons" bson:"seasons"`
}

type DriverSeason struct {
	Year string `json:"year" bson:"year"`
	Team string `json:"team" bson:"team"`
}

var DriverType *graphql.Object
var DriverSeasonType *graphql.Object

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

	DriverSeasonType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "DriverSeason",
		Description: "Dummy object to join the season type and the team type inside the DriverType",
		Fields: graphql.Fields{
			"season": &graphql.Field{
				Name:        "season",
				Type:        SeasonType,
				Description: "The season this driver has driven",
			},
			"team": &graphql.Field{
				Name:        "team",
				Type:        TeamType,
				Description: "The team this driver has driven for",
			},
		},
	})
}
