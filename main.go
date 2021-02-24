package main

import (
	"log"

	//"github.com/gorilla/mux"
	"github.com/sploitfun/idesk/db"
)

func main() {

	//Connect to DB
	client, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	/* Init router
	r := mux.NewRouter()

	r.HandleFunc("/candidates", getCandidates).Methods("GET")
	r.HandleFunc("/candidate/{id}", getCandidate).Methods("GET")
	r.HandleFunc("/candidate", createCandidate).Methods("POST")
	r.HandleFunc("/candidate/{id}", updateCandidate).Methods("PUT")
	r.HandleFunc("/candidate/{id}", deleteCandidate).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))*/
}
