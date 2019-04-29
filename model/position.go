package model

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Position represents a F1 Position
type Position struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Driver Driver             `json:"driver" bson:"driver"`
	GP     GP                 `json:"gp" bson:"gp"`
	Number int                `json:"number" bson:"number"`
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
			"gp": &graphql.Field{
				Name:        "gp",
				Type:        GPType,
				Description: "The gp the driver got this position",
			},
			"number": &graphql.Field{
				Name:        "number",
				Type:        graphql.Int,
				Description: "The position itself",
			},
		},
	})

}
