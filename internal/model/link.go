package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID            primitive.ObjectID `bson:"_id"`
	OriginLink    string             `bson:"origin_link"`
	ShortenedLink string             `bson:"shortened_link"`
	Code          string             `bson:"code"`
	VisitorAmount int64              `bson:"visitor_amount"`
	Status        string             `bson:"status"`
	ExpiredAt     *time.Time         `bson:"expired_at"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}
