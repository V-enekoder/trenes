package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func getDatabaseConnection(ctx context.Context) (neo4j.SessionWithContext, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("URI")
	username := os.Getenv("NEO4J_USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Print(uri, "\n", username, "\n", password, "\n")
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, fmt.Errorf("error al crear el driver: %w", err)
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, fmt.Errorf("error de conexión: %w", err)
	}

	session := driver.NewSession(ctx, neo4j.SessionConfig{})

	return session, nil
}
