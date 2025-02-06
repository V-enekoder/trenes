package station

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindOptimalRoadController(c *gin.Context) {
	startIDStr := c.Param("start_id") // Corregido: usar nombres de parámetros distintos
	startID, err := strconv.ParseInt(startIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de inicio inválido"})
		return
	}

	endIDStr := c.Param("end_id") // Corregido: usar nombres de parámetros distintos
	endID, err := strconv.ParseInt(endIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de fin inválido"})
		return
	}

	path, weight, err := FindOptimalRoadService(c, startID, endID) // Pasa el contexto
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"path": path, "weight": weight})
}

// Controlador para crear una estación
func CreateStationController(c *gin.Context) {
	var estacion Station
	if err := c.ShouldBindJSON(&estacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := CreateStationService(c, estacion) // Pasa el contexto del gin
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, estacion)
}

// Controlador para obtener una estación por ID
func GetStationByIDController(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	estacion, err := GetStationByIDService(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if estacion == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Estación no encontrada"})
		return
	}

	c.JSON(http.StatusOK, estacion)
}

// Controlador para obtener todas las estaciones
func GetAllStationsController(c *gin.Context) {
	estaciones, err := GetAllStationsService(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, estaciones)
}

func UpdateStationController(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var estacion Station
	if err := c.ShouldBindJSON(&estacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asigna el ID del parámetro a la estación
	estacion.ID = id

	err = UpdateStationService(c, estacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, estacion)
}

// Controlador para eliminar una estación por ID
func DeleteStationController(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = DeleteStationService(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estación eliminada correctamente"})
}
