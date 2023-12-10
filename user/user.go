package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Role string

const (
	Master    Role = "master"
	Admin     Role = "admin"
	Blogger   Role = "blogger"
	Merchant  Role = "merchant"
	Publisher Role = "publisher"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"` // Store hashed password
	Role      Role               `bson:"role"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
