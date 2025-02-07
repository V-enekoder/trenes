package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/V-enekoder/trenes/config"
	station "github.com/V-enekoder/trenes/src/stations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	ctx := context.Background()
	session, err := config.GetDatabaseConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close(ctx)

	// TODO: Execute code
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	station.SetupRoutes(r)

	r.Run()

	// test(session, ctx)
}

func test(session neo4j.SessionWithContext, ctx context.Context) {
	query := `MATCH (e:Estacion)
		RETURN e.Id AS id, e.name AS name, e.line AS line, e.typestation AS typestation, e.system AS system`
	result, err := session.Run(ctx, query, nil)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	// Itera sobre los resultados
	for result.Next(ctx) {
		record := result.Record()
		id := record.Values[0]
		name := record.Values[1]
		line := record.Values[2]
		typestation := record.Values[3]
		system := record.Values[4]

		fmt.Printf("ID: %v, Nombre: %v, Línea: %v, Tipo: %v, Sistema: %v\n", id, name, line, typestation, system)
	}

	// Verifica si hubo algún error durante la iteración
	if err = result.Err(); err != nil {
		log.Fatalf("Error en los resultados: %v", err)
	}
}
