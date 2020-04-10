package mongo

import "go.mongodb.org/mongo-driver/bson"

func (mc MongoClient) InsertSomethingToTest() error {
	c := mc.database.Collection("test_collection")

	_, err := c.InsertOne(nil, bson.M{"test_field": "test_document"})

	return err
}
