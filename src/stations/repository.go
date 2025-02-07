package station

import (
	"context"
	"fmt"
	"log"

	"github.com/V-enekoder/trenes/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func createStationRepository(ctx context.Context, estacion Station) error {
	session := config.SESSION

	cypher := `CREATE (e:Estacion {Id: $id, name: $name, line: $line, typestation: $typestation, system: $system})`
	_, err := session.Run(ctx, cypher,
		map[string]interface{}{
			"id":          estacion.ID,
			"name":        estacion.Name,
			"line":        estacion.Line,
			"typestation": estacion.Typestation,
			"system":      estacion.System,
		})
	if err != nil {
		return fmt.Errorf("error creando la estación: %w", err)
	}

	return nil
}

func getStationByIdRepository(ctx context.Context, id int64) (*Station, error) {
	session := config.SESSION

	cypher := `MATCH (e:Estacion {Id: $id}) RETURN e`
	result, err := session.Run(ctx, cypher, map[string]interface{}{"id": id})
	if err != nil {
		return nil, fmt.Errorf("error ejecutando la consulta: %w", err)
	}

	if result.Next(ctx) {
		record := result.Record()
		node := record.Values[0].(neo4j.Node)
		props := node.GetProperties() // Llama a la función Props()
		estacion := &Station{
			ID:          props["Id"].(int64), // Accede a las propiedades del mapa
			Name:        props["name"].(string),
			Line:        props["line"].(int64),
			Typestation: props["typestation"].(string),
			System:      props["system"].(string),
		}
		return estacion, nil
	}

	return nil, nil // No se encontró la estación
}
func getAllStationsRepository(ctx context.Context) ([]*Station, error) {
	session := config.SESSION

	cypher := `MATCH (e:Estacion) RETURN e`
	result, err := session.Run(ctx, cypher, nil)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando la consulta: %w", err)
	}

	var estaciones []*Station
	for result.Next(ctx) {
		record := result.Record()
		for _, v := range record.Values {
			log.Println(v)
		}
		node := record.Values[0].(neo4j.Node)
		props := node.GetProperties()
		estacion := &Station{
			ID:          props["Id"].(int64),
			Name:        props["name"].(string),
			Line:        props["line"].(int64),
			Typestation: props["typestation"].(string),
			System:      props["system"].(string),
		}
		estaciones = append(estaciones, estacion)
	}

	return estaciones, nil
}

// Actualizar una estación
func UpdateStationRepository(ctx context.Context, estacion Station) error {
	session := config.SESSION

	cypher := `MATCH (e:Estacion {Id: $id})
	SET e.name = $name, e.line = $line, e.typestation = $typestation, e.system = $system`
	_, err := session.Run(ctx, cypher,
		map[string]interface{}{
			"id":          estacion.ID,
			"name":        estacion.Name,
			"line":        estacion.Line,
			"typestation": estacion.Typestation,
			"system":      estacion.System,
		})
	if err != nil {
		return fmt.Errorf("error actualizando la estación: %w", err)
	}

	return nil
}

// Eliminar una estación por su ID
func DeleteStationRepository(ctx context.Context, id int64) error {
	session := config.SESSION

	cypher := `MATCH (e:Estacion {Id: $id}) DETACH DELETE e`
	_, err := session.Run(ctx, cypher, map[string]interface{}{"id": id})
	if err != nil {
		return fmt.Errorf("error eliminando la estación: %w", err)
	}

	return nil
}

func findOptimalPathRepository(ctx context.Context, startID, endID int64) (op OptimalPath, err error) {
	session := config.SESSION

	cypher := `MATCH (start:Estacion {Id: $startID}) MATCH (end:Estacion {Id: $endID})
		CALL apoc.algo.dijkstra(start, end, 'CONNECTS_TO', 'distance')
		YIELD path, weight
		RETURN REDUCE(s = [], n IN nodes(path) | s + n.name) AS nombres_estaciones, weight
	`
	result, err := session.Run(ctx, cypher, map[string]interface{}{"startID": startID, "endID": endID})
	if err != nil {
		err = fmt.Errorf("error ejecutando Dijkstra: %w", err)
		return op, err
	}

	if result.Next(ctx) {
		record := result.Record()
		op.Path = record.Values[0].([]interface{}) // path es un arreglo con el nombre de los nodos
		for _, v := range op.Path {
			log.Println(v)
		}
		op.Weight = record.Values[1].(float64)
		op.Time = op.Weight / 90

		return op, err
	}
	err = fmt.Errorf("no se encontró ruta entre las estaciones %d y %d", startID, endID)
	return op, err
}
