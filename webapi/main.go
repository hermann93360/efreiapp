package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"
)

type whoami struct {
	Name  string
	Title string
	State string
}

func main() {
	request1()
}

func whoAmI(response http.ResponseWriter, r *http.Request) {
	who := []whoami{
		whoami{Name: "Efrei Paris",
			Title: "DevOps and Continous Deployment",
			State: "FR",
		},
	}

	json.NewEncoder(response).Encode(who)

	fmt.Println("Endpoint Hit", who)
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Web API!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "EfreiParis"

	fmt.Fprintf(response, "A little bit about me...")
	fmt.Println("Endpoint Hit: ", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoamiHandler)
	http.HandleFunc("/error", error)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func whoamiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Team Name: Hermann talla Oussama Franck Team")
}

func error(w http.ResponseWriter, r *http.Request) {
    logMessage := fmt.Sprintf("ERROR: Something went wrong at %s", time.Now().Format(time.RFC3339))
	log.Println(logMessage)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}