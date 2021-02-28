package main

import (
	"flag"
	"fmt"

	"github.com/cnaize/oat/database"
	"github.com/cnaize/oat/server"
)

func main() {
	serverPort := flag.Uint("port", 8080, "server port")
	sourceType := flag.String("source", "json", "source type")
	flag.Parse()

	questionsDB, err := database.NewQuestions((database.SourceType)(*sourceType))
	if err != nil {
		fmt.Printf("db initialization failed: %+v", err)
		return
	}
	db := database.DB{
		Questions: questionsDB,
	}
	if err := server.NewServer(&db).Run(*serverPort); err != nil {
		fmt.Printf("server run failed: %+v\n", err)
	}
}
