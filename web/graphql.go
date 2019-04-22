package web

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
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
