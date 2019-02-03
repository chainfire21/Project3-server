package mongo

import(
	"context"
	"log"
	"time"
	"os"
	"strings"
	"net/http"
	"encoding/json"
	"Project3-server/typeform"

	// "github.com/Jeffail/gabs"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/mongodb/mongo-go-driver/bson"

)

var myClient = &http.Client{Timeout: 10 * time.Second}

func connectServer() (coll *mongo.Collection, CancelFunc func()){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, os.Getenv("MONGODB_URI"))
	// client, err := mongo.Connect(ctx, "mongodb://heroku_w02f0l1k:30fj40p12gho8osfmp81qd1oq7@ds213755.mlab.com:13755/heroku_w02f0l1k")

	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("MONGODB_DB")).Collection("users")
	// collection := client.Database("heroku_w02f0l1k").Collection("users")


	return collection, cancel
}

func GetMatches(e string){
	// user := GetUser(e)
	collection, cancel := connectServer()
	defer cancel()

	var result UserModel
	cur, err := collection.Find(context.Background(), bson.D{bson.E{Key:"usertype",Value:"coach"}})
	if err != nil { log.Fatal(err) }
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
	   raw, err := cur.DecodeBytes()
	   if err != nil { log.Fatal(err) }
	   // do something with elem....
	   bson.Unmarshal(raw,&result)
	   log.Println(result)
	   log.Println(result.Results)
	}
	if err := cur.Err(); err != nil {
			log.Fatal(err)
	}
}

func UpdateUser(e string) {
	coll, cancel :=  connectServer()
	defer cancel()
	var r Results
	surveys := typeform.GetSurveyDataClient()
	children, _ := surveys.S("items").Children()
	for _, child := range children {
		exists := child.ExistsP("answers.email")
		if exists == true{
			email := child.Path("answers.email").String()
			trimE1 := strings.Replace(email, "[", "", -1)
			trimE2 := strings.Replace(trimE1, "]", "", -1)
			trimE3 := strings.Replace(trimE2,"\"","",-1)
			if trimE3 == e{
				loc := child.Path("answers.text").String()
				trimL1 := strings.Replace(loc, "[", "", -1)
				trimL2 := strings.Replace(trimL1, "]", "", -1)
				trimL3 := strings.Replace(trimL2,"\"","",-1)
				r.Location = trimL3
				choice := child.Path("answers.choice.label").Data()
				switch choice := choice.(type) {
				case []interface{}:
					for index, value := range choice {
						if s, ok := value.(string); ok { 
							if index == 0{
								r.NewClients = s
							}
							if index == 1{
								r.Types = []string{s}
							}
							
						}

					}
				}
				choices := child.Path("answers.choices.labels").Data()
				switch choices := choices.(type) {
				case []interface{}:
					for index, value := range choices {
						switch value := value.(type) {
						case []interface{}:
							var str []string
							for _, value2 := range value{
								if s, ok := value2.(string); ok { 
									if index == 0{
										str = append(str,s)
										r.GorOne = str
									}
									if index == 1{
										str = append(str,s)
										r.Virtual = str
									}
									if index == 2{
										str = append(str,s)
										r.Gender = str
									}
									if index == 3{
										str = append(str,s)
										r.Topics = str
									}
									if index == 4{
										str = append(str,s)										
										r.Traits = str
									}
									
								}
							}
						}
						
					}
				}
				log.Println(r)
				jsonR,err := json.Marshal(r)
				if err != nil{
					log.Fatal(err)
				}
				log.Println(string(jsonR))
				emailDoc := bson.M{"email": e}
				resInt := bson.M{"$set":bson.M{"results":r}}
				updated, err := coll.UpdateOne(
					context.Background(),
					emailDoc,
					resInt,
					
				)
				log.Println(updated)
			}
		}
	}

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
	filter := bson.M{"email": email}
	log.Println(filter)
	err := collection.FindOne(context.Background(),filter).Decode(&result)
	if err != nil{
		log.Fatal(err)
	}
	return result
}