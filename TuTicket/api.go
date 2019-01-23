package main

import (
	"net/http"
	"strconv"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Obtiene todos los eventos
func getAllEventos(c *gin.Context) {
	//Obtiene la variable de la conexión
	db := c.MustGet("DB").(*gorm.DB)

	//Encuentra todos los eventos
	var eventos []EventoModel
	db.Find(&eventos)

	if len(eventos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡Eventos no encontrados!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": eventos})
}

//Obtiene un solo evento
func getSingleEvento(c *gin.Context) {
	//Obtiene la variable de la conexión
	db := c.MustGet("DB").(*gorm.DB)

	//Encuentra el evento con el id que se envió como parámetro
	var evento EventoModel
  	eventoID := c.Param("evento_id")
	db.First(&evento, eventoID)

	if evento.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡Evento no encontrado!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": evento})
}