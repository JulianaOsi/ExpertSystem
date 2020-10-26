package main

import (
	"context"
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
	connPool, err := store.New(conf)
	if err != nil {
		log.Fatalf("create a connection pool failed: %v", err)
	}

	symptom := &store.Symptom{
		Id:   1,
		Name: "Боль в горле",
	}
	if err := connPool.CreateOrUpdateSymptom(context.Background(), symptom); err != nil {
		log.Fatalf("create or update sympom failed: %v", err)
	}
}
