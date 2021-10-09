package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User      primitive.ObjectID `json:"user" bson:"user,omitempty"`
	Caption   string             `json:"caption,omitempty" bson:"caption,omitempty"`
	Image_url string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Timestamp time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

type PostsResponse struct {
	Posts     []Post `json:"posts"`
	Total     int64  `json:"total"`
	Page      int64  `json:"page"`
	Last_page int64  `json:"last_page"`
	Limit     int64  `json:"limit"`
}
