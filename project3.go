package main

import (
	"net/http"
	"os"

	"Project3-server/mongo"
	// "Project3-server/typeform"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/mongodb/mongo-go-driver/bson"
)

func checkProd(p string) (prt string){
	if p == ""{
		return ":1323"
	}
	return ":"+p
}

func main() {
	// Echo instance
	e := echo.New()

	// matches := mongo.GetMatches("f@gmail.com")
	// log.Println(matches)
	// jsonString, _ := json.Marshal(matches)
	// log.Println(jsonString)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/user/:email", func(c echo.Context) error{
		return c.JSON(200,mongo.GetUser(c.Param("email")))
	})

	e.GET("/matches/:email", func(c echo.Context) error{
		mongo.UpdateUser(c.Param("email"))
		matches := mongo.GetMatches(c.Param("email"))
		return c.JSON(200, matches)
	})

	e.POST("/newuser", func(c echo.Context) error{
		m := mongo.UserModel{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		return c.JSON(200, mongo.AddUser(m))	
	})
	port := os.Getenv("PORT")

	// Start server
	e.Logger.Fatal(e.Start(checkProd(port)))

}
