package controllers

import (  
    //"encoding/json"
    "encoding/base64"
    //"log"
    "bytes"
    "time"
    //"fmt"
    "net/http"
    "github.com/labstack/echo"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"

    "video-for-kids/models"
    "video-for-kids/utils"
)

type (
    // UserController represents the controller for operating on the User resource
    UserController struct {
        session *mgo.Session
    }
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(s *mgo.Session) *UserController {
    return &UserController{s}
}

type H map[string]interface{}

func(uc UserController) GetUser(c echo.Context) error {
    u := models.User{}
    id := c.Param("id")
    //check ID
    if !bson.IsObjectIdHex(id) {
        return c.JSON(http.StatusNotFound,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusNotFound,
                    "title": "Not Found",
                    "detail": "Cannot find resource for specified ID.",
                },
            })
    }
    // Grab id
    oid := bson.ObjectIdHex(id)
    // Fetch user
    if err := uc.session.DB("go_rest_tutorial").C("users").FindId(oid).One(&u); err != nil {
        return c.JSON(http.StatusNotFound,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusNotFound,
                    "title": "Not Found",
                    "detail": "Cannot find resource for specified ID.",
                },
            })
    }
    //building JSON data
    resp := bson.M{
            "data": bson.M{
                "id": u.Id,
                "fullname": u.Fullname,
                "username": u.Username,
            },
        }
    //return JSON if no errors
    return c.JSON(http.StatusOK, resp)
}

//Create user
func (uc UserController) CreateUser(c echo.Context) error{
    // Stub an user to be populated from the body
    u := models.User{}
    if err := c.Bind(&u); err != nil {
        return err
    }

    //MongoDB query pipeline
    pipeline := bson.M{
        "$or": []interface{}{
            bson.M{"email": u.Email},
            bson.M{"username": u.Username},
        },
    }

    // Check if user exists
    count, err := uc.session.DB("go_rest_tutorial").C("users").Find(pipeline).Count()

    if err != nil {
        return c.JSON(http.StatusInternalServerError,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusInternalServerError,
                    "title": "Internal Server Errord",
                    "detail": "The server encountered an unexpected condition which prevented it from fulfilling the request.",
                },
            })
    }

    if count > 0 {
        return c.JSON(http.StatusConflict,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusConflict,
                    "title": "Resources Already Exists",
                    "detail": "The request could not be completed due to a duplicate with the current state of the resource.",
                },
            })
    }

    //generate Authorization key (base64)
    var str bytes.Buffer
    str.WriteString(u.Email)
    str.WriteString(":")
    str.WriteString(u.Password)
    u.AuthorizationKey = base64.StdEncoding.EncodeToString([]byte(str.String()))

    //Hash password
    u.Password, err = utils.HashPassword(u.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusInternalServerError,
                    "title": "Internal Server Errord",
                    "detail": "The server encountered an unexpected condition which prevented it from fulfilling the request.",
                },
            })
    }

    // Add an Id
    u.Id = bson.NewObjectId()

    //set default roles
    u.Roles = []string{"contributor", "admin"}

    //set default active
    u.Active = true

    //generate access token and set expiry
    u.AccessToken, err = utils.NewUUID()
    u.AccessTokenExpiry = time.Now().AddDate(0, 0, 30)

    // Write the user to mongo
    if err := uc.session.DB("go_rest_tutorial").C("users").Insert(u); err != nil {
        return c.JSON(http.StatusInternalServerError,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusInternalServerError,
                    "title": "Internal Server Errord",
                    "detail": "The server encountered an unexpected condition which prevented it from fulfilling the request.",
                },
            })
    }
    //building JSON data
    data := H{
            "data": H{
                "id": u.Id,
                "fullname": u.Fullname,
                "username": u.Username,
                "roles": u.Roles,
                "access_token": u.AccessToken,
                "access_token_expiry": u.AccessTokenExpiry,
            },
        }
    return c.JSON(http.StatusCreated, data)
}

//Update user
func (uc UserController) UpdateUser(c echo.Context) error{
    // Stub an user to be populated from the body
    u := models.User{}

    //Get the ID
    id := c.Param("id")
    // Grab id
    oid := bson.ObjectIdHex(id)

    if err := c.Bind(&u); err != nil {
        return err
    }

    // Verify id is ObjectId, otherwise bail
    //check ID
    if !bson.IsObjectIdHex(id) {
        return c.JSON(http.StatusNotFound,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusNotFound,
                    "title": "Not Found",
                    "detail": "Cannot find resource for specified ID.",
                },
            })
    }
    /*
    //generate Authorization key (base64)
    var str bytes.Buffer
    str.WriteString(u.Email)
    str.WriteString(":")
    str.WriteString(u.Password)
    u.AuthorizationKey = base64.StdEncoding.EncodeToString([]byte(str.String()))
    
    //Hash password
    u.Password, err = utils.HashPassword(u.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError,     
            models.Response{
                Errors: models.Error{
                    Status: http.StatusInternalServerError,
                    Title: "Internal Server Error",
                    Detail: "The server encountered an unexpected condition which prevented it from fulfilling the request",
                },
            })
    }

    */
    //set default roles
    u.Roles = []string{"contributor", "admin"}

    //set default active
    u.Active = true

    // Write the user to mongo
    change := bson.M{"$set": bson.M{"fullname": u.Fullname, "roles": u.Roles}}  
    if err := uc.session.DB("go_rest_tutorial").C("users").UpdateId(oid, change); err != nil {
        return c.JSON(http.StatusNotFound,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusNotFound,
                    "title": "Not Found",
                    "detail": "Cannot find resource for specified ID.",
                },
            })
    }

    //json response
    data := H{
            "data": H{
                "id": u.Id,
                "fullname": u.Fullname,
                "username": u.Username,
                "roles": u.Roles,
                "access_token": u.AccessToken,
                "access_token_expiry": u.AccessTokenExpiry,
            },
        }
    return c.JSON(http.StatusOK, data)
}

//remove user
func (uc UserController) DeleteUser(c echo.Context) error{
    id := c.Param("id")

    // Verify id is ObjectId, otherwise bail
    //check ID
    if !bson.IsObjectIdHex(id) {
        return c.JSON(http.StatusNotFound,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusNotFound,
                    "title": "Not Found",
                    "detail": "Cannot find resource for specified ID.",
                },
            })
    }

    // Grab id
    oid := bson.ObjectIdHex(id)

    // Remove user
    if err := uc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
        return c.JSON(http.StatusNotFound,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusNotFound,
                    "title": "Not Found",
                    "detail": "Cannot find resource for specified ID.",
                },
            })
    }

    // Write status
    resp := bson.M{
        "data": bson.M{
            "id": id, 
            "message":"User has been deleted",
            },
        }
    return c.JSON(http.StatusOK, resp)
}

//Get user videos
func(uc UserController) GetUserVideos(c echo.Context) error {
    //u := models.User{}
    id := c.Param("id")
    //check ID
    if !bson.IsObjectIdHex(id) {
        return c.JSON(http.StatusNotFound,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusNotFound,
                    "title": "Not Found",
                    "detail": "Cannot find resource for specified ID.",
                },
            })
    }
    // Grab id
    oid := bson.ObjectIdHex(id)

    //MongoDB query pipeline
    pipeline := []bson.M{
            bson.M{ "$match" : bson.M{ "_id" : oid } },
            bson.M{
              "$lookup":
                bson.M{
                  "from": "videos",
                  "localField": "_id",
                  "foreignField": "submitted_by",
                  "as": "videos",
                },
           },
           bson.M{ "$unwind": 
               bson.M{
                   "path": "$videos",
                   "preserveNullAndEmptyArrays": true,
               },
           },
           bson.M{
             "$group": bson.M{
                 "_id": "$_id",
                 "username": bson.M{ "$first":"$username" },
                 "fullname": bson.M{ "$first":"$fullname" },
                 "videos": bson.M{"$push": "$videos"},
                 },
           },
           bson.M{
             "$project": bson.M{
                 "_id": 1,
                 "username": 1,
                 "fullname": 1,
                 "videos": bson.M{
                        "_id": 1,
                        "title": 1,
                        "url": 1,
                        "active": 1,
                        "submitted_date": 1,
                        "categories": 1,
                     },
             },
           },
    }

    // Fetch user
    resp := bson.M{}
    if err := uc.session.DB("go_rest_tutorial").C("users").Pipe(pipeline).One(&resp); err != nil {
        return c.JSON(http.StatusNotFound,     
            bson.M{
                "errors": bson.M{
                    "status": http.StatusNotFound,
                    "title": "Not Found",
                    "detail": "Cannot find resource for specified ID.",
                },
            })
            
    } 
    //return JSON if no errors
    return c.JSON(http.StatusOK, resp)
}

//authentication
func Auth(username, password string) (bool, error) {
    if username == "joe" && password == "secret" {
        return true, nil
    }
    return false, nil
}