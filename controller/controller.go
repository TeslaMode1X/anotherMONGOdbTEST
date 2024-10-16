package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongoAPI/model"
	"net/http"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const collectionName = "watchlist"

// VERY IMPORTANT
var collection *mongo.Collection

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connection to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)

	// collection reference
	fmt.Println("Collection Reference is ready")
}

// MONGODB helpers

// get all movies from mongo
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie primitive.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}

// insert 1 record
func insertMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("One movie was inserted into DB", inserted)
}

// update 1 record
func updateMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println("modified count: ", res.ModifiedCount, "\n", res)
}

// delete 1 record
func deleteMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count", deleteCount)
}

// delete all records
func deleteAllMovies() {
	delRes, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All records are deleted: ", delRes, "\n", delRes.DeletedCount)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	deleteAllMovies()

	json.NewEncoder(w).Encode("All movies are deleted")
}
