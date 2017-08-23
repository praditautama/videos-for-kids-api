package main

import (
	"fmt"
	"time"
	//"net"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"crypto/tls"
)

type Game struct {
	Winner       string    `bson:"winner"`
	OfficialGame bool      `bson:"official_game"`
	Location     string    `bson:"location"`
	StartTime    time.Time `bson:"start"`
	EndTime      time.Time `bson:"end"`
	Players      []Player  `bson:"players"`
}

type Player struct {
	Name   string    `bson:"name"`
	Decks  [2]string `bson:"decks"`
	Points uint8     `bson:"points"`
	Place  uint8     `bson:"place"`
}

func NewPlayer(name, firstDeck, secondDeck string, points, place uint8) Player {
	return Player{
		Name:   name,
		Decks:  [2]string{firstDeck, secondDeck},
		Points: points,
		Place:  place,
	}
}

var isDropMe = true

func main() {
	Host := []string{
		"localhost:27017",
		// replica set addrs...
	}
	const (
		Username   = "tutorial"
		Password   = "1234"
		Database   = "go_rest_tutorial"
		Collection = "games"
	)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
		//Username: Username,
		//Password: Password,
		//Database: Database,
		//DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
		// 	return tls.Dial("tcp", addr.String(), &tls.Config{})
		//},
	})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	game := Game{
		Winner:       "Dave",
		OfficialGame: true,
		Location:     "Austin",
		StartTime:    time.Date(2015, time.February, 12, 04, 11, 0, 0, time.UTC),
		EndTime:      time.Now(),
		Players: []Player{
			NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
			NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
			NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
			NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
		},
	}

	if isDropMe {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection
	c := session.DB(Database).C(Collection)

	// Insert
	if err := c.Insert(game); err != nil {
		panic(err)
	}

	// Find and Count
	player := "Dave"
	gamesWon, err := c.Find(bson.M{"winner": player}).Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has won %d games.\n", player, gamesWon)

	// Find One (with Projection)
	var result Game
	err = c.Find(bson.M{"winner": player, "location": "Austin"}).Select(bson.M{"official_game": 1}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Is game in Austin Official?", result.OfficialGame)

	// Find All
	var games []Game
	err = c.Find(nil).Sort("-start").All(&games)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of Games", len(games))

	// Update
	newPlayer := "John"
	selector := bson.M{"winner": player}
	updator := bson.M{"$set": bson.M{"winner": newPlayer}}
	if err := c.Update(selector, updator); err != nil {
		panic(err)
	}

	// Update All
	info, err := c.UpdateAll(selector, updator)
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated", info.Updated)

	// Remove
	
	//info, err = c.RemoveAll(bson.M{"winner": newPlayer})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Removed", info.Removed)
}
