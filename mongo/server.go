package mongo

import(
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"

)

func AddUser(u *UserModel) (val interface{}){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	// client, err := mongo.Connect(ctx, "mongodb://heroku_w02f0l1k:30fj40p12gho8osfmp81qd1oq7@ds213755.mlab.com:13755/heroku_w02f0l1k")
	if err != nil{
		log.Fatal(err)
	}
	
	// collection := client.Database("heroku_w02f0l1k").Collection("numbers")
	collection := client.Database("testing").Collection("numbers")
	log.Println("HEY u")
	log.Println(u)
	res, err := collection.InsertOne(ctx, u)
	id := res.InsertedID
	return id
}

func GetUser(email string) (res UserModel){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, "mongodb://heroku_w02f0l1k:30fj40p12gho8osfmp81qd1oq7@ds213755.mlab.com:13755/heroku_w02f0l1k")
	if err != nil{
		log.Fatal(err)
	}
	
	collection := client.Database("heroku_w02f0l1k").Collection("numbers")
	// collection := client.Database("testing").Collection("numbers")

	var result UserModel
	filter := bson.D{{"email", email}}
	log.Println(filter)
	error := collection.FindOne(context.Background(),filter).Decode(&result)
	if err != nil{
		log.Fatal(error)
	}
	return result
}