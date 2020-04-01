package mongo

import "go.mongodb.org/mongo-driver/bson"

func (mc MongoClient) InsertSomethingToTest() error {
	c := mc.database.Collection("test")

	_, err := c.InsertOne(nil, bson.M{"test": "test mongo 2"})

	return err
}
