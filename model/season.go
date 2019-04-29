package model

import (
	"github.com/graphql-go/graphql"
)

// Season represents a F1 season
type Season struct {
	Year   string  `json:"year" bson:"_id"`
	Tracks []Track `json:"tracks" bson:"tracks"`
}

var SeasonType *graphql.Object

func init() {
	SeasonType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Season",
		Description: "A Formula 1 Season",
		Fields: graphql.Fields{
			"year": &graphql.Field{
				Name:        "year",
				Type:        graphql.String,
				Description: "Season's year",
			},
			"tracks": &graphql.Field{
				Name:        "tracks",
				Type:        graphql.NewList(TrackType),
				Description: "Seasons's tracks",
			},
		},
	})
}
