package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func LaunchServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", welcomeHandler).Methods(http.MethodGet)
	r.HandleFunc("/test", symptomHandler).Methods(http.MethodGet)
	r.HandleFunc("/test/{symptom}", testStartHandler).Methods(http.MethodGet)
	r.HandleFunc("/test/{symptom}/{knowledge}", testHandler).Methods(http.MethodGet)
	r.HandleFunc("/knowledge", knowledgeSectionHandler).Methods(http.MethodGet)
	r.HandleFunc("/knowledge/{symptom}/add", addKnowledgeHandler).Methods(http.MethodPost)
	r.HandleFunc("/knowledge/{symptom}/", knowledgeHandler).Methods(http.MethodGet)

	fmt.Printf("Starting server at localhost:8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
