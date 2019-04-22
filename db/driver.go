package db

import (
	"errors"

	"github.com/eltonjr/graphql-server-exercise/model"
)

type DriverDao struct {
}

func NewDriverDao() *DriverDao {
	return &DriverDao{}
}

func (dd *DriverDao) GetDrivers() ([]*model.Driver, error) {
	return nil, errors.New("TODO: method not implemented")
}

func (dd *DriverDao) GetSingleDriver(id string) (*model.Driver, error) {
	return nil, errors.New("TODO: method not implemented")
}

func (dd *DriverDao) CreateDriver(*model.Driver) (*model.Driver, error) {
	return nil, errors.New("TODO: method not implemented")
}

func (dd *DriverDao) UpdateDriver(*model.Driver) (*model.Driver, error) {
	return nil, errors.New("TODO: method not implemented")
}

func (dd *DriverDao) DeleteDriver(id string) (*model.Driver, error) {
	return nil, errors.New("TODO: method not implemented")
}
