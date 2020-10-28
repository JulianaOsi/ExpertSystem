package main

import (
	"context"
	"fmt"
	"log"

	"expertSystem/src/store"
)

func main() {
	conf := &store.ConfigDB{
		Host: "localhost",
		Port: "5432",
		Name: "expertSystem",
		User: "postgres",
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
}
