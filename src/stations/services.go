package station

import (
	"context"
	"fmt"
)

// Servicio para crear una estación
func CreateStationService(ctx context.Context, estacion Station) error {
	err := createStationRepository(ctx, estacion)
	if err != nil {
		return fmt.Errorf("error en el servicio de creación de estación: %w", err)
	}
	return nil
}

// Servicio para obtener una estación por ID
func GetStationByIDService(ctx context.Context, id int64) (*Station, error) {
	estacion, err := getStationByIdRepository(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error en el servicio de obtención de estación por ID: %w", err)
	}
	return estacion, nil
}

// Servicio para obtener todas las estaciones
func GetAllStationsService(ctx context.Context) ([]*Station, error) {
	estaciones, err := getAllStationsRepository(ctx)
	if err != nil {
		return nil, fmt.Errorf("error en el servicio de obtención de todas las estaciones: %w", err)
	}
	return estaciones, nil
}

// Servicio para actualizar una estación
func UpdateStationService(ctx context.Context, estacion Station) error {
	err := UpdateStationRepository(ctx, estacion)
	if err != nil {
		return fmt.Errorf("error en el servicio de actualización de estación: %w", err)
	}
	return nil
}

// Servicio para eliminar una estación por ID
func DeleteStationService(ctx context.Context, id int64) error {
	err := DeleteStationRepository(ctx, id)
	if err != nil {
		return fmt.Errorf("error en el servicio de eliminación de estación: %w", err)
	}
	return nil
}

func FindOptimalPathService(ctx context.Context, startID, endID int64) (OptimalPath, error) {
	return findOptimalPathRepository(ctx, startID, endID)
}
