package mongo

import(
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"

)

func addUser(u *UserModel) (val interface{}, name string){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	// client, err := mongo.Connect(ctx, "mongodb://heroku_w02f0l1k:30fj40p12gho8osfmp81qd1oq7@ds213755.mlab.com:13755/heroku_w02f0l1k")
	if err != nil{
		log.Fatal(err)
	}
	// log.Println(client)	
	collection := client.Database("testing").Collection("numbers")
	res, err := collection.InsertOne(ctx, u)
	id := res.InsertedID
	return id, u.Name
}