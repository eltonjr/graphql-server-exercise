package db

import (
	"context"
	"fmt"

	"github.com/eltonjr/graphql-server-exercise/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const driversCollection = "drivers"

func GetDrivers(skip, limit int) ([]*model.Driver, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	opt := options.Find()
	opt.SetSkip(int64(skip))
	opt.SetLimit(int64(limit))

	cursor, err := db.Collection(driversCollection).Find(context.Background(), bson.M{}, opt)
	if err != nil {
		return nil, fmt.Errorf("could not find documents in '%s' collection: %v", driversCollection, err)
	}

	drivers := []*model.Driver{}
	for cursor.Next(context.Background()) {
		curr := &model.Driver{}
		err := cursor.Decode(curr)
		if err != nil {
			return nil, fmt.Errorf("error decoding driver: %v", err)
		}
		drivers = append(drivers, curr)
	}

	return drivers, nil
}

func GetSingleDriver(id primitive.ObjectID) (*model.Driver, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	var driver model.Driver

	filter := bson.M{"_id": id}

	err = db.Collection(driversCollection).FindOne(context.Background(), filter).Decode(&driver)
	if err != nil {
		return nil, fmt.Errorf("could not find driver with id '%s': %v", id, err)
	}

	return &driver, nil
}

func CreateDriver(driver *model.Driver) (*model.Driver, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	driver.ID = primitive.NewObjectID()

	inserted, err := db.Collection(driversCollection).InsertOne(context.Background(), driver)
	if err != nil {
		return nil, fmt.Errorf("could not insert driver: %v", err)
	}

	driver.ID = inserted.InsertedID.(primitive.ObjectID)

	return driver, nil
}

func UpdateDriver(driver *model.Driver) (*model.Driver, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": driver.ID}

	_, err = db.Collection(driversCollection).UpdateOne(context.Background(), filter, driver)
	if err != nil {
		return nil, fmt.Errorf("could not update driver: %v", err)
	}

	return driver, nil
}

func DeleteDriver(id primitive.ObjectID) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}

	_, err = db.Collection(driversCollection).DeleteOne(context.Background(), filter)

	return err
}
