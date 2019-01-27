package main

import (
	"net/http"
	"context"
	"time"
	"log"
	// "fmt"
    // "io/ioutil"
    "os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mongodb/mongo-go-driver/mongo"
	// "github.com/mongodb/mongo-go-driver/bson"
)

func main() {
	// Echo instance
	e := echo.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, "mongodb://heroku_w02f0l1k:30fj40p12gho8osfmp81qd1oq7@ds213755.mlab.com:13755/heroku_w02f0l1k")
	if err != nil{
		log.Fatal(err)
	}
	log.Println(client)	
	// collection := client.Database("testing").Collection("numbers")
	// res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	// id := res.InsertedID
	// log.Println(id)
	
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	// response, err := http.Get("https://api.typeform.com/forms/GA1xBQ/responses")
	// if err != nil {
    //     fmt.Printf("%s", err)
    //     os.Exit(1)
    // } else {
    //     defer response.Body.Close()
    //     contents, err := ioutil.ReadAll(response.Body)
    //     if err != nil {
    //         fmt.Printf("%s", err)
    //         os.Exit(1)
    //     }
    //     fmt.Printf("%s\n", string(contents))
    // }
	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/newuser", func(c echo.Context) error{
		return c.JSON(http.StatusOK, "{'response':'good'}")
	})

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// Start server
	e.Logger.Fatal(e.Start(":"+port))
}
