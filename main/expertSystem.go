package main

import (
	"context"
	"fmt"
	"log"

	"expertSystem/src/store"
)

func main() {
	conf := &store.ConfigDB{
		Host:     "ekb.urai.ru",
		Port:     "9811",
		Name:     "ExpertSystem",
		User:     "postgres",
		Password: "******",
	}
	s, err := store.New(conf)
	if err != nil {
		log.Fatalf("create a connection pool failed: %v", err)
	}

	symptoms, err := s.GetAllSymptoms(context.Background())
	if err != nil {
		log.Fatalf("get all symptoms failed: %v", err)
	}
	for _, s := range symptoms {
		fmt.Println(s.Name)
	}

	questions, err := s.GetQuestionsBySymptomId(context.Background(), "1")
	if err != nil {
		log.Fatalf("get questions by symptomId failed: %v", err)
	}
	for _, q := range questions {
		fmt.Println(q.Text)
	}
}
