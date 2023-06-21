package database

const (
	DatabaseName   = "go_rest_api"
	CollectionName = "users"
)

type Database struct {
	Username string
	Password string
	Host     string
}

var db = Database{
	Username: "admin",
	Password: "admin",
	Host:     "localhost",
}

func SetDatabase(username, password string) {
	db.Username = username
	db.Password = password
}

func Login(username, password string) bool {
	if username == db.Username && password == db.Password {
		return true
	}
	return false
}

func GetDatabase() Database {

	return db
}
