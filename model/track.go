package model

import (
	"github.com/graphql-go/graphql"
)

// Track represents a F1 track
type Track struct {
	ID      string   `json:"id" bson:"_id"`
	Name    string   `json:"name" bson:"name"`
	Country string   `json:"country" bson:"country"`
	Years   []string `json:"years" bson:"years"`
	GPs     []GP     `json:"gps" bson:"-"`
}

// GP represents a F1 GP - a track race in a given year
type GP struct {
	Season  string     `json:"season" bson:"_id"`
	TrackID string     `json:"-" bson:"trackId"`
	Track   Track      `json:"track" bson:"-"`
	Results []Position `json:"results" bson:"results"`
}

var TrackType *graphql.Object
var GPType *graphql.Object

func init() {
	TrackType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Track",
		Description: "A Formula 1 Track",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Name:        "id",
				Type:        graphql.String,
				Description: "Track's id",
			},
			"name": &graphql.Field{
				Name:        "name",
				Type:        graphql.String,
				Description: "Track's name",
			},
			"country": &graphql.Field{
				Name:        "country",
				Type:        graphql.String,
				Description: "Track's country",
			},
			"years": &graphql.Field{
				Name:        "years",
				Type:        graphql.NewList(graphql.String),
				Description: "Track's years",
			},
			"gps": &graphql.Field{
				Name:        "gps",
				Type:        graphql.NewList(GPType),
				Description: "Track's gps",
			},
		},
	})

	GPType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GP",
		Description: "An instance of the track in a given year",
		Fields: graphql.Fields{
			"season": &graphql.Field{
				Name:        "season",
				Type:        graphql.String,
				Description: "The season this GP was",
			},
			"track": &graphql.Field{
				Name:        "track",
				Type:        TrackType,
				Description: "The full track object for this GP",
			},
			"results": &graphql.Field{
				Name:        "results",
				Type:        graphql.NewList(PositionType),
				Description: "The GP's results",
			},
		},
	})
}
