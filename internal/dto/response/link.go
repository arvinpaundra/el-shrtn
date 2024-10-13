package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	CreatedLink struct {
		ID            primitive.ObjectID `json:"id"`
		ShortenedLink string             `json:"shortened_link"`
		ExpiredAt     *string            `json:"expired_at"`
	}
)
