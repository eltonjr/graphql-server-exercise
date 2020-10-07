package model

import (
	"time"

	"github.com/graphql-go/graphql"
)

// Team represents a F1 scuderia
type Team struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	CreatedAt *time.Time   `json:"createdAt"`
	Seasons   []TeamSeason `json:"seasons" bson:"seasons"`
}

type TeamSeason struct {
	Year   string `json:"year" bson:"year"`
	Driver string `json:"driver" bson:"driver"`
}

var TeamType *graphql.Object
var TeamSeasonType *graphql.Object

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

	TeamSeasonType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "TeamSeason",
		Description: "Dummy object to join the season type and the driver type inside the TeamType",
		Fields: graphql.Fields{
			"season": &graphql.Field{
				Name:        "season",
				Type:        SeasonType,
				Description: "The season this driver has driven",
			},
			"driver": &graphql.Field{
				Name:        "driver",
				Type:        DriverType,
				Description: "The driver this team had",
			},
		},
	})
}
