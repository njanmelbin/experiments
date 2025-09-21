package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

func main() {

	// Replace these values with your actual connection data
	host := "localhost"
	port := 5431
	user := "postgres"
	password := "mysecretpassword" // The password you set for the container
	dbname := "postgres"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	fmt.Println("Max open connections:", db.Stats().OpenConnections)
	fmt.Println("Max idle connections:", db.Stats().Idle)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			var n int
			err := db.QueryRow("SELECT 1").Scan(&n)
			if err != nil {
				log.Printf("Query error: %v", err)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("Open connections:", db.Stats().OpenConnections)
	fmt.Println("Idle connections:", db.Stats().Idle)

}
