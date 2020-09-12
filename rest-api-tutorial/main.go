package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type article struct {
	Title string `json:"title"`
	Desc  string `json:"desc`
}

type articles []article

func allarticles(w http.ResponseWriter, r *http.Request) {
	articles := articles{
		article{
			Title: "hello world from mux",
			Desc:  "hello world desc",
		},
	}
	json.NewEncoder(w).Encode(articles)

}

func articlesPostRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post end point hit")

}

func handleRequest() {
	// http.HandleFunc("/", homepage)

	myRouter := mux.NewRouter().StrictSlash(true)
	/** for gorilla mux */
	myRouter.HandleFunc("/articles", allarticles).Methods("GET")
	myRouter.HandleFunc("/articles", articlesPostRequest).Methods("POST")
	/** gorilla mux */

	/** default go router
	// http.HandleFunc("/articles", allarticles)
	**/

	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	handleRequest()
}
