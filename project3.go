package main

import (
	"net/http"
	"context"
	"time"
	"log"
	"fmt"
    "io/ioutil"
    "os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
)

func main() {
	// Echo instance
	e := echo.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	// client, err := mongo.Connect(ctx, "MLAB STUFF")
	if err != nil{
		log.Fatal(err)
	}	
	collection := client.Database("testing").Collection("numbers")
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	log.Println(id)
	
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	response, err := http.Get("https://api.typeform.com/forms/GA1xBQ/responses")
	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }
	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.POST("/newuser", func(c echo.Context) error{
		return c.JSON(http.StatusOK, "{'response':'good'}")
	})


	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}