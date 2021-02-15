package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `json:"title"`
	Year        int                `json:"year"`
	Length      int                `json:"length"`
	Description string             `json:"description"`
	Director    string             `json:"director"`
	Genre       string             `json:"genre"`
	Rating      int                `json:"rating"`
	AgeRating   int                `json:"agerating"`
}