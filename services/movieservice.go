package services

import (
	"context"
	"fmt"
	"github.com/CH02z/movielibrary/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var err error

type MovieService struct {
	dbclient *mongo.Client
	mongoDb  *mongo.Database
}

func newService(connectString string, dbName string) MovieService {
	fmt.Printf("trying to establish a connection to a mongo database server...\n")
	clientOpts := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("database Connectino established.. ping the database to check functionality\n")
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Printf("failed to ping the mongo db server: %s\n", err)
	}

	fmt.Printf("configure usage of database with name '%s'\n", dbName)
	return MovieService{dbclient: client, mongoDb: client.Database(dbName)}
}

func (s MovieService) GetAllMovies() ([]models.Movie, error) {

	var moviesBSON []bson.M
	var movies []models.Movie

	fmt.Printf("select all movies from collection\n")

	collection := s.mongoDb.Collection("movies")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("fatal in mongo repo with connection")
		log.Fatal(err)
	}

	for cursor.Next(context.Background()) {
		var movie bson.M
		err = cursor.Decode(&movie)
		if err == nil {
			moviesBSON = append(moviesBSON, movie)
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("found '%d' movies in collection\n", len(movies))

	// convert []primitve.m to array of structs (movie.Mongomovie)
	for _, bsonelement := range moviesBSON {
		var mov models.Movie
		bsonBytes, _ := bson.Marshal(bsonelement)
		bson.Unmarshal(bsonBytes, &mov)
		movies = append(movies, mov)
	}
	return movies, err
}

func (s MovieService) GetOneMovie(movieID string) (models.Movie, error) {
	var movie models.Movie
	collection := s.mongoDb.Collection("movies")
	objID, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": objID}
	err := collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		fmt.Println(err)
	}
	return movie, err
}

func (s MovieService) CreateMovie(newMovie models.Movie) (interface{}, error) {
	movie := &models.Movie{ID: primitive.NewObjectIDFromTimestamp(time.Now()),
		Title:       newMovie.Title,
		Description: newMovie.Description,
		Genre:       newMovie.Genre,
		Year:        newMovie.Year,
		Length:      newMovie.Length,
		Director:    newMovie.Director,
		Rating:      newMovie.Rating,
		AgeRating:   newMovie.AgeRating}
	collection := s.mongoDb.Collection("movies")
	insertResult, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult.InsertedID, err
}

func (s MovieService) DeleteMovie(movieID string) error {
	collection := s.mongoDb.Collection("movies")
	objID, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": objID}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

