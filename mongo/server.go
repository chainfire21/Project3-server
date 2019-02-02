package mongo

import(
	"context"
	"log"
	"time"
	"os"
	// "bytes"
	"net/http"
	// "encoding/json"
	"Project3-server/typeform"

	// "github.com/Jeffail/gabs"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"

)

var myClient = &http.Client{Timeout: 10 * time.Second}

func connectServer() (coll *mongo.Collection, CancelFunc func()){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("MONGODB_DB")).Collection("users")
	// collection := client.Database("testing").Collection("numbers")

	return collection, cancel
}

func GetMatches(e string){
	matches := typeform.GetSurveyData()
	log.Println(matches)
}

func AddUser(u UserModel) (val interface{}){
	collection, cancel := connectServer()
	defer cancel()

	_, err := collection.InsertOne(context.Background(), u)
	if err != nil{
		log.Fatal(err)
		return err
	}
	return "User Created"
}

func GetUser(email string) (res UserModel){
	collection, cancel := connectServer()
	defer cancel()

	var result UserModel
	filter := bson.D{{"email", email}}
	log.Println(filter)
	err := collection.FindOne(context.Background(),filter).Decode(&result)
	if err != nil{
		log.Fatal(err)
	}
	return result
}