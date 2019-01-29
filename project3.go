package main

import (
	"net/http"
	"log"
	"fmt"
    "io/ioutil"
	"os"
	"bytes"
	"encoding/json"

	// "Project3-server/mongo"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/mongodb/mongo-go-driver/bson"
)

func getSurveyData(userType string){
	var prettyJSON bytes.Buffer
	httpClient := &http.Client{}
	if userType == "user"{
		fmt.Println("HI DETECTED")
	}
	
	req, err := http.NewRequest("GET", "https://api.typeform.com/forms/GA1xBQ/responses",nil)
	req.Header.Add("Authorization", "Bearer Gr8o49DXvMTTVnDaCNjz86mS2kE283snRA4S25ULogmk")
	response, err := httpClient.Do(req)
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
		error := json.Indent(&prettyJSON, contents, "", "\t")
		if error != nil {
			log.Println("JSON parse error: ", error)
			return
		}
		fmt.Printf("%s\n", string(prettyJSON.Bytes()))
	}
}

func checkProd(p string) (prt string){
	if p == ""{
		return ":1323"
	}
	return ":"+p
}

func main() {
	// Echo instance
	e := echo.New()

	
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.POST("/newuser", func(c echo.Context) error{
		// log.Println("IN POST")
		// log.Println(c)
		// var u mongo.UserModel
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		return c.JSON(200, m)	
	})
	port := os.Getenv("PORT")

	// Start server
	e.Logger.Fatal(e.Start(checkProd(port)))

}
