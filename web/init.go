package web

import (
	"fmt"
	"graph-analyzer/api/config"
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func InitDBConnection() neo4j.Driver {
	driver, err := neo4j.NewDriver(
		fmt.Sprintf("%s:%s", config.GetEnvVariable("NEO4J_HOST"), config.GetEnvVariable("NEO4J_PORT")),
		neo4j.BasicAuth(
			config.GetEnvVariable("NEO4J_USER"),
			config.GetEnvVariable("NEO4J_PASSWORD"),
			config.GetEnvVariable("NEO4J_REALM"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Try to verify database connection 10 times
	for i := 1; i <= 10; i++ {
		log.Printf("Trying to verify database connection, try #%d of %d\n", i, 10)
		err = driver.VerifyConnectivity()
		if err == nil {
			break
		}

		if i == 10 {
			log.Fatalln("Connection to database could not be established")
		}

		log.Printf("Failed to establish database connection, retrying in %d seconds\n", 10)

		time.Sleep(10 * time.Second)
	}

	log.Println("Database connection established")

	return driver
}
