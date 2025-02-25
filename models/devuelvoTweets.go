package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweets es la estructura donde devuelvo los tweets*/
type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`

	Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
}
