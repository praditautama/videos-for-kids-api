package main 

import (  
    // Standard library packages
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "github.com/TobiEiss/goMiddlewareChain"
    "github.com/TobiEiss/goMiddlewareChain/templates"
    "gopkg.in/mgo.v2"

    "video-for-kids/controllers"
    "video-for-kids/middlewares"
)

func main() {
    // Instantiate a new router
    r := httprouter.New()

    // Get a UserController instance
    uc := controllers.NewUserController(getSession())

    // Get a user resource
    r.GET("/user/:id", goMiddlewareChain.RestrictedRequestChainHandlerWithResponseCheck(true, middlewares.Auth, templates.JSONResponseHandler, middlewares.Logging, uc.GetUser))

    // Create a new user
    r.POST("/user", goMiddlewareChain.RestrictedRequestChainHandler(middlewares.Auth, templates.JSONResponseHandler, middlewares.Logging, uc.CreateUser))

    // Remove an existing user
    r.DELETE("/user/:id", goMiddlewareChain.RestrictedRequestChainHandler(middlewares.Auth, templates.JSONResponseHandler, middlewares.Logging, uc.RemoveUser))

    // Fire up the server
    http.ListenAndServe("localhost:3000", r)
}

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