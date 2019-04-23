package web

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/eltonjr/graphql-server-exercise/db"
	"github.com/eltonjr/graphql-server-exercise/model"
)

func (router *Router) GetDriversResolver(p graphql.ResolveParams) (interface{}, error) {
	return db.GetDrivers(0, 10) // TODO move this to somewhere else
}

func (router *Router) GetSingleDriverResolver(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"]
	if id == nil {
		return nil, errors.New("missing id")
	}

	idStr, ok := id.(string)
	if !ok {
		return nil, fmt.Errorf("id must be a string: %v", id)
	}

	if idStr == "" {
		return nil, errors.New("id must not be empty")
	}

	idHex, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid driver id: %v", err)
	}

	return db.GetSingleDriver(idHex)
}

func (router *Router) CreateDriverResolver(p graphql.ResolveParams) (interface{}, error) {
	name, _ := p.Args["name"].(string)
	country, _ := p.Args["country"].(string)

	driver := &model.Driver{
		Name:    name,
		Country: country,
	}

	return db.CreateDriver(driver)
}

func (router *Router) UpdateDriverResolver(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	name, _ := p.Args["name"].(string)
	country, _ := p.Args["country"].(string)

	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid drive id '%s': %v", id, err)
	}

	driver := &model.Driver{
		ID:      hexID,
		Name:    name,
		Country: country,
	}

	return db.UpdateDriver(driver)
}

func (router *Router) DeleteDriverResolver(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid driver id: %v", err)
	}

	err = db.DeleteDriver(idHex)
	return nil, err
}
