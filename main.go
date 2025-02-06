package main

import (
	"context"
	"fmt"
	"log"

	"github.com/V-enekoder/trenes/config"
)

func main() {
	ctx := context.Background()

	session, err := config.GetDatabaseConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (e:Estacion) RETURN e.Id AS id, e.name AS name, e.line AS line, e.typestation AS typestation, e.system AS system", nil)
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
