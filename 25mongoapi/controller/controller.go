package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BlackDevil559/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://scrapernerd:Cg24122002@cluster0.mvhrh.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// Connect with mongoDB
func init(){
	clientOption:=options.Client().ApplyURI(connectionstring)
	client,err:=mongo.Connect(context.TODO(),clientOption)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection Sucess")
	collection=client.Database(dbName).Collection(colName)
	fmt.Println("Collection Intance is Ready")
}

// Insert 1 Record
func insertOneMovie(movie model.Netflix){
	inserted,err:=collection.InsertOne(context.Background(),movie)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Movie is inserted with ID",inserted.InsertedID)
}

// Update 1 Record
func updateOneRecord(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}
	result,err:=collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Modified count",result.ModifiedCount)
}

// Delete 1 Record
func deleteOneRecord(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	result,err:=collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Modified count",result.DeletedCount)
}

// Delete Many Record
func deleteManyRecord(){
	filter:=bson.D{{}}
	result,err:=collection.DeleteMany(context.Background(),filter,nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Modified count",result.DeletedCount)
}

// Get all records from MongoDB
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return movies
}

func GetAllMovies(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all movies")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(getAllMovies())
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	fmt.Println("Creating Movie")
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
	fmt.Println("Moive added succesfully")
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	params:=mux.Vars(r)
	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode("Mark watched successfully")
	fmt.Println("Movie marked")
}

func DeleteOneRecord(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params:=mux.Vars(r)
	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode("Movie deleted")
	fmt.Println("Movie deleted")
}

func DeleteManyRecord(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	deleteManyRecord()
	json.NewEncoder(w).Encode("all Movie deleted")
	fmt.Println("all Movie deleted")
}