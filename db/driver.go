package db

import (
	"github.com/eltonjr/graphql-server-exercise/model"
)

type DriverDao struct {
	dummyData []*model.Driver
}

func NewDriverDao() *DriverDao {
	return &DriverDao{
		dummyData: []*model.Driver{
			&model.Driver{
				ID:      "1",
				Name:    "Kimi Raikkonen",
				Country: "Finland",
			},
			&model.Driver{
				ID:      "2",
				Name:    "Romain Grosjean",
				Country: "France",
			},
			&model.Driver{
				ID:      "3",
				Name:    "Nico Hulkenberg",
				Country: "Germany",
			},
			&model.Driver{
				ID:      "4",
				Name:    "Daniel Riccardo",
				Country: "Australia",
			},
		},
	}
}

func (dd *DriverDao) GetDrivers() ([]*model.Driver, error) {
	return dd.dummyData, nil
}

func (dd *DriverDao) GetSingleDriver(id string) (*model.Driver, error) {
	return dd.dummyData[0], nil
}

func (dd *DriverDao) CreateDriver(*model.Driver) (*model.Driver, error) {
	return dd.dummyData[1], nil
}

func (dd *DriverDao) UpdateDriver(*model.Driver) (*model.Driver, error) {
	return dd.dummyData[2], nil
}

func (dd *DriverDao) DeleteDriver(id string) (*model.Driver, error) {
	return dd.dummyData[3], nil
}
