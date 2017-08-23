package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)



/*
users
- userid 
- username
- fullname
- email
- password_hash
- roles
- active
- access_token
*/

type (
	// User represents the structure of our resource
	User struct {
		Id                bson.ObjectId `json:"id" bson:"_id"`
		Username          string        `json:"username" bson:"username"`
		Fullname          string        `json:"fullname" bson:"fullname"`
		Email             string        `json:"email,omitempty" bson:"email"`
		Password          string        `json:"password,omitempty" bson:"password"`
		Roles             []string      `json:"roles,omitempty" bson:"roles"`
		Active            bool          `json:"active,omitempty" bson:"active"`
		AuthorizationKey  string        `json:"authorization_key,omitempty" bson:"authorization_key"`
		AccessToken       string        `json:"access_token,omitempty" bson:"access_token"`
		AccessTokenExpiry time.Time     `json:"access_token_expiry,omitempty" bson:"access_token_expiry"`
	}
)