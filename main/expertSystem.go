package main

import (
	"expertSystem/src/config"
	migrations "expertSystem/src/migration"
	"expertSystem/src/store"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	conf := config.ReadConfig()

	err := migrations.UpMigrations(conf)
	if err != nil {
		log.Fatalf("failed to update migrations: %v\n", err)
	}

	if err := store.InitDB(conf.DB); err != nil {
		log.Fatalf("failed to create store: %v\n", err)
	}

	//server.LaunchServer()
	/*
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
	*/
}
