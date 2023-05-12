package mongo

import "go.mongodb.org/mongo-driver/mongo"

type Logging struct {
	db *mongo.Database
}

func NewLogging(db *mongo.Database) *Logging {
	return &Logging{
		db: db,
	}
}

func (l Logging) Insert() {
	//TODO implement me
	panic("implement me")
}
