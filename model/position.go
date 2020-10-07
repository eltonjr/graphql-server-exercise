package model

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Position represents a F1 Position
type Position struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	DriverID string             `json:"-" bson:"driverId"`
	Driver   Driver             `json:"driver" bson:"-"`
	Number   int                `json:"number" bson:"number"`
}

var PositionType *graphql.Object

func init() {
	PositionType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Position",
		Description: "A GP position for a single Driver",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Name:        "id",
				Type:        graphql.String,
				Description: "Position's id",
			},
			"driver": &graphql.Field{
				Name:        "driver",
				Type:        DriverType,
				Description: "The driver who got this position",
			},
			"number": &graphql.Field{
				Name:        "number",
				Type:        graphql.Int,
				Description: "The position itself",
			},
		},
	})

}
