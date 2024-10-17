package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email" json:"email"`
}

type Order struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
	Burger Burger
}

type Burger struct {
	Cheese  int `bson:"cheese" json:"cheese"`
	Meat    int `bson:"meat" json:"meat"`
	Lettuce int `bson:"lettuce" json:"lettuce"`
	Tomato  int `bson:"tomato" json:"tomato"`
}
