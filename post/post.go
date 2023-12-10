package post

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Type string
type ParagraphType string
type Status string
type Locale string
type Category string

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	Locale      Locale             `json:"locale" bson:"locale"`
	Status      Status             `json:"status" bson:"status"`
	Image       string             `json:"image" bson:"image"`
	Title       string             `json:"title" bson:"title"`
	Desc        string             `json:"desc" bson:"desc"`
	URLTitle    string             `json:"urlTitle" bson:"urlTitle"`
	Date        time.Time          `json:"date" bson:"date"`
	Href        string             `json:"href" bson:"href"`
	Categories  []Category         `json:"categories" bson:"categories"`
	Comments    []Comments         `json:"comments" bson:"comments"`
	Paragraphs  []Paragraph        `json:"paragraphs" bson:"paragraphs"`
	Views       int                `json:"views" bson:"views"`
	ReadingTime int                `json:"readingTime" bson:"readingTime"`
	Like        Like               `json:"like" bson:"like"`
	AuthorID    primitive.ObjectID `json:"authorId" bson:"authorId"`
	Type        Type               `json:"type" bson:"type"`
}

type Paragraph struct {
	ParagraphType string `json:"type" bson:"paragraphType"`
	Text          string `json:"text" bson:"text"`
	Img           struct {
		Src string `json:"src" bson:"src"`
	} `json:"img" bson:"img"`
}

type Comments struct {
	UserID   string     `json:"userId" bson:"userId"`
	Date     string     `json:"date" bson:"date"`
	Content  string     `json:"content" bson:"content"`
	ParentID string     `json:"parentId" bson:"parentId"`
	Like     Like       `json:"like" bson:"like"`
	Children []Comments `json:"children" bson:"children"`
}

type Like struct {
	Count int `json:"count"`
}
