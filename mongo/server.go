package mongo

import(
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"

)

func connectServer() (coll *mongo.Collection, CancelFunc func()){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, "mongodb://heroku_w02f0l1k:30fj40p12gho8osfmp81qd1oq7@ds213755.mlab.com:13755/heroku_w02f0l1k")
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("heroku_w02f0l1k").Collection("numbers")
	// collection := client.Database("testing").Collection("numbers")

	return collection, cancel
}

func AddUser(u UserModel) (val interface{}){
	collection, cancel := connectServer()
	defer cancel()

	res, err := collection.InsertOne(context.Background(), u)
	if err != nil{
		log.Fatal(err)
	}
	id := res.InsertedID
	return id
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