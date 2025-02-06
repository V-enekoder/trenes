package station

import (
	"context"
	"fmt"

	"github.com/V-enekoder/trenes/config"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func createStationRepository(ctx context.Context, estacion Station) error {
	session, err := config.GetDatabaseConnection(ctx)
	if err != nil {
		return fmt.Errorf("error obteniendo la conexión a la base de datos: %w", err)
	}
	defer session.Close(ctx)

	_, err = session.Run(ctx, "CREATE (e:Estacion {Id: $id, name: $name, line: $line, typestation: $typestation, system: $system})",
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
	session, err := config.GetDatabaseConnection(ctx)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo la conexión: %w", err)
	}
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (e:Estacion {Id: $id}) RETURN e", map[string]interface{}{"id": id})
	if err != nil {
		return nil, fmt.Errorf("error ejecutando la consulta: %w", err)
	}

	if result.Next(ctx) {
		record := result.Record()
		node := record.Values[0].(neo4j.Node)
		props := node.Props() // Llama a la función Props()
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
	session, err := config.GetDatabaseConnection(ctx)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo la conexión: %w", err)
	}
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (e:Estacion) RETURN e", nil)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando la consulta: %w", err)
	}

	var estaciones []*Station
	for result.Next(ctx) {
		record := result.Record()
		node := record.Values[0].(neo4j.Node)
		props := node.Props()
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
	session, err := config.GetDatabaseConnection(ctx)
	if err != nil {
		return fmt.Errorf("error obteniendo la conexión: %w", err)
	}
	defer session.Close(ctx)

	_, err = session.Run(ctx, "MATCH (e:Estacion {Id: $id}) SET e.name = $name, e.line = $line, e.typestation = $typestation, e.system = $system",
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
	session, err := config.GetDatabaseConnection(ctx)
	if err != nil {
		return fmt.Errorf("error obteniendo la conexión: %w", err)
	}
	defer session.Close(ctx)

	_, err = session.Run(ctx, "MATCH (e:Estacion {Id: $id}) DETACH DELETE e", map[string]interface{}{"id": id})
	if err != nil {
		return fmt.Errorf("error eliminando la estación: %w", err)
	}

	return nil
}
