package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	//"video-for-kids/models/user"
	//"video-for-kids/models/video"
)


/*
category
- categoryid
- name
- added_date
- added_by
- videos
*/

type (
	// User represents the structure of our resource
	Category struct {
		Id           bson.ObjectId `json:"id" bson:"_id"`
		Name         string        `json:"name" bson:"name"`
		AddedDate    time.Time     `json:"added_date" bson:"added_date"`
		AddedBy      bson.ObjectId          `json:"added_by" bson:"added_by"`
	}
)