package main

import (
	"net/http"
	"fmt"
	"regexp"
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"

	"video-for-kids/controllers"
    //"video-for-kids/middleware"
)

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://tutorial:12345@localhost:27017/go_rest_tutorial")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }

    // Deliver session
    return s
}

func Login(c echo.Context) error {
    return c.JSON(http.StatusOK, "4896bbed-2415-4eb5-a595-570ebc9c6316")
}

func main() {
	
    // New echo
	e := echo.New()
	//remove trailing slash
	e.Pre(middleware.RemoveTrailingSlash())
	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(func (username, password string, c echo.Context) (bool, error) {
		
		path := c.Path()
		login, _ := regexp.MatchString("login", path)
		fmt.Println(c.Path(), login)
		if login {
			//by pass
			//if username == "hello@praditautama.com" && password == "1234" {
				return true, nil
			//}
			return true, nil
		} else {
			//use userid:token
			if username == "1111" && password == "4896bbed-2415-4eb5-a595-570ebc9c6316" {
				return true, nil
			}
			return false, nil
		}
		
	}))

	//default route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to API v1!")
	})

	//Users
	// Get a UserController instance
    uc := controllers.NewUserController(getSession())
	e.POST("/users", uc.CreateUser)
	e.GET("/users/:id", uc.GetUser)
	e.PATCH("/users/:id", uc.UpdateUser)
	e.DELETE("/users/:id", uc.DeleteUser)
	e.GET("/users/:id/videos", uc.GetUserVideos)

	e.POST("/login", Login)



	e.Logger.Fatal(e.Start(":1323"))
}
