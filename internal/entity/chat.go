package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Chat stores user's or group chat info.
type Chat struct {
	Id        primitive.ObjectID `bson:"_id"`
	ChatId    int                `bson:"chat_id"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
	Country   string             `bson:"country"`
}

func (c *Chat) MarshalBSON() ([]byte, error) {
	c.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	type tmp Chat
	return bson.Marshal((*tmp)(c))
}
