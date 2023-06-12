package localhost

import (
	"fmt"
	"log"
	"net/http"
)

// here we will make the web server and the api calls for displaying the data and the alerts
// it will be a webinterface for the user to see the data and the alerts

func host() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// start the web server

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
