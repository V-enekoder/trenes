package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	session, err := getDatabaseConnection(ctx)
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
	/*estaciones, err := session.BeginTransaction(ctx, func(tx neo4j.ExplicitTransaction) (interface{}, error) {
		query := "MATCH (e:Estacion) RETURN e.Id AS id, e.name AS name, e.line AS line, e.typestation AS typestation, e.system AS system"
		result, err := tx.Run(ctx, query, nil) // Pasa el contexto a tx.Run
		if err != nil {
			return nil, err
		}

		var estaciones []map[string]interface{}
		for result.Next(ctx) { // Pasa el contexto a result.Next
			record := result.Record()
			estaciones = append(estaciones, map[string]interface{}{
				"id":          record.Values[0],
				"name":        record.Values[1],
				"line":        record.Values[2],
				"typestation": record.Values[3],
				"system":      record.Values[4],
			})
		}

		// Verifica si hubo errores durante la iteración
		if err = result.Err(); err != nil {
			return nil, err
		}

		return estaciones, nil
	})

	if err != nil {
		log.Fatalf("Error al consultar estaciones: %v", err)
	}

	fmt.Println("Estaciones:")
	for _, e := range estaciones.([]map[string]interface{}) {
		fmt.Printf("ID: %v, Nombre: %v, Línea: %v, Tipo: %v, Sistema: %v\n", e["id"], e["name"], e["line"], e["typestation"], e["system"])
	}*/
}
