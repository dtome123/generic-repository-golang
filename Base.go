package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Base struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt int64              `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt int64              `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt int64              `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
