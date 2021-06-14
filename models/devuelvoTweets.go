package models

import (
	"time"

	"go.mongodb.org/driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* DevuelvoTweets es la estructura con la que devolverermos los tweets*/
type DevueltoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}