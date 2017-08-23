package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)


/*
api_logs
- logid
- access_date
- endpoint_url
- json_request
- json_response
*/

type (
	// User represents the structure of our resource
	Logs struct {
		Id             bson.ObjectId `json:"id" bson:"_id"`
		AccessDate     time.Time     `json:"access_date" bson:"access_date"`
		EndpointUrl    string        `json:"endpoint_url" bson:"endpoint_url"`
		JsonRequest    string        `json:"json_request" bson:"json_request"`
		JsonResponse   string        `json:"json_response" bson:"json_response"`
	}
)