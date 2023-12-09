package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	CreateUser(ctx context.Context, user User) (*mongo.InsertOneResult, error)
	GetUser(ctx context.Context, id string) (User, error)
	UpdateUser(ctx context.Context, id string, user User) (*mongo.UpdateResult, error)
	DeleteUser(ctx context.Context, id string) (*mongo.DeleteResult, error)
}

type Repository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *Repository {
	collection := db.Collection("users")
	return &Repository{collection: collection}
}

func (r *Repository) CreateUser(ctx context.Context, user User) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, user)
}

func (r *Repository) GetUser(ctx context.Context, id string) (User, error) {
	var user User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (r *Repository) UpdateUser(ctx context.Context, id string, user User) (*mongo.UpdateResult, error) {
	return r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{{"$set", user}},
	)
}

func (r *Repository) DeleteUser(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(ctx, bson.M{"_id": id})
}
