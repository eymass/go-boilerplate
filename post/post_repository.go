package post

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IPostRepository interface {
	Create(ctx context.Context, post Post) (*mongo.InsertOneResult, error)
	Read(ctx context.Context, id primitive.ObjectID) (Post, error)
	Update(ctx context.Context, id primitive.ObjectID, post Post) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetByCategories(ctx context.Context, categories []string, page, pageSize int) ([]Post, error)
}

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}

func (r *Repository) Create(ctx context.Context, post Post) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, post)
}

func (r *Repository) Read(ctx context.Context, id primitive.ObjectID) (Post, error) {
	var post Post
	err := r.collection.FindOne(ctx, Post{ID: id}).Decode(&post)
	return post, err
}

func (r *Repository) GetByHref(ctx context.Context, href string) (Post, error) {
	var post Post
	err := r.collection.FindOne(ctx, Post{Href: href}).Decode(&post)
	return post, err
}

func (r *Repository) GetByCategories(ctx context.Context, categories []Category, page, pageSize int) ([]Post, error) {
	var posts []Post
	findOptions := options.Find().SetSkip(int64(pageSize * (page - 1))).SetLimit(int64(pageSize))

	filter := bson.M{"categories": bson.M{"$in": categories}}
	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *Repository) Update(ctx context.Context, id primitive.ObjectID, post Post) (*mongo.UpdateResult, error) {
	return r.collection.UpdateOne(ctx, Post{ID: id}, post)
}

func (r *Repository) Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(ctx, Post{ID: id})
}
