package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Admin     Role = "Admin"
	Blogger   Role = "Blogger"
	Merchant  Role = "Merchant"
	Publisher Role = "Publisher"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"` // Store hashed password
	Role     Role               `bson:"role"`
}
