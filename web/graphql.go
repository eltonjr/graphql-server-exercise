package web

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"

	"github.com/eltonjr/graphql-server-exercise/model"
)

func (router *Router) GetDriversResolver(p graphql.ResolveParams) (interface{}, error) {
	return router.driverDao.GetDrivers()
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

	return router.driverDao.GetSingleDriver(idStr)
}

func (router *Router) CreateDriverResolver(p graphql.ResolveParams) (interface{}, error) {
	name, _ := p.Args["name"].(string)
	country, _ := p.Args["country"].(string)

	driver := &model.Driver{
		Name:    name,
		Country: country,
	}

	return router.driverDao.CreateDriver(driver)
}

func (router *Router) UpdateDriverResolver(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	name, _ := p.Args["name"].(string)
	country, _ := p.Args["country"].(string)

	driver := &model.Driver{
		ID:      id,
		Name:    name,
		Country: country,
	}

	return router.driverDao.UpdateDriver(driver)
}

func (router *Router) DeleteDriverResolver(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	return router.driverDao.DeleteDriver(id)
}
