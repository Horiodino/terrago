package api

// here going to define the mongodb and all the data going to be saved here

import ()

// here going to define the mongodb and all the data going to be saved here
// first of all we need to define the database name

const (
	// DatabaseName is the name of the database
	DatabaseName = "go_rest_api"
	// CollectionName is the name of the collection
	CollectionName = "users"
)

// here going to define the mongodb and all the data going to be saved here
// adn we will also added authentication here so we can access the database for the web application

// craeting a authentication username and password for the database
// here we are going to create a struct for the database

type Database struct {
	Username string
	Password string
	Host     string
}

// now set the username and password for the database
// by default the username and password is empty
// then we will ask the user to enter the username and password

var db = Database{
	Username: "admin",
	Password: "admin",
	Host:     "localhost",
}

// now we need to create a function to set new username and password
// so we can access the database

func SetDatabase(username, password string) {
	db.Username = username
	db.Password = password
}

// to save the datda in the mongodb we need to use json format and we also need to
// define the struct for the data which we are going to save in the mongodb

// that will be done in metrics.go file