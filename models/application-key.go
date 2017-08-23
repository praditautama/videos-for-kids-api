package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)


/*
application_keys
- api_key
- provision_date
- application_name
- email
- active
*/

type (
	// User represents the structure of our resource
	ApplicationKey struct {
		Id              bson.ObjectId `json:"id" bson:"_id"`
		ApiKey          string        `json:"api_key" bson:"api_key"`
		ProvisionDate   time.Time     `json:"provision_date" bson:"provision_date"`
		ApplicationName string        `json:"application_name" bson:"application_name"`
		Email           string        `json:"email" bson:"email"`
		Active          bool          `json:"active" bson:"active"`
	}
)