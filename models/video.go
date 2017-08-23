package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	//"video-for-kids/models/user"
	//"video-for-kids/models/category"
)

/*
videos
- videoid
- title
- url
- provider
- submitted_by
- submitted_date
- categories
- active
*/

type (
	// User represents the structure of our resource
	Video struct {
		Id            bson.ObjectId `json:"id" bson:"_id"`
		Title         string        `json:"title" bson:"title"`
		Url           string        `json:"url" bson:"url"`
		Provider      string        `json:"provider" bson:"provider"`
		SubmitedBy    bson.ObjectId `json:"submitted_by" bson:"submitted_by"`
		SubmitedDate  time.Time     `json:"submitted_date" bson:"submitted_date"`
		Categories    []bson.ObjectId    `json:"categories" bson:"categories"`
		Active        bool          `json:"active" bson:"active"`
	}
)