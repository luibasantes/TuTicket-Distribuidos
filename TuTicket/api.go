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

//Obtiene todos los asientos de un evento
func getAsientos(c *gin.Context) {
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

	//Encuentra los asientos del evento
	var asientos []AsientoModel
	db.Model(&evento).Related(&asientos);
	
	if len(asientos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡No hay asientos!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": asientos})
}

//Compra un boleto
func buyBoleto(c *gin.Context) {
	//Obtiene la variable de la conexión
	db := c.MustGet("DB").(*gorm.DB)

	//Encuentra el asiento con el id que se envió en el POST
	var asiento AsientoModel
	asientoID, _ := strconv.Atoi(c.PostForm("asientoID"))
	db.First(&asiento, asientoID)

	if asiento.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡Asiento no encontrado!"})
		return
	}

	//Valida que el asiento esté disponible
	if asiento.Disponible == false {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusBadRequest, "message": "¡El asiento no está disponible!"})
		return
	}

	//Encuentra el usuario con el id que se envió en el POST
	var usuario UsuarioModel
  	usuarioID, _ := strconv.Atoi(c.PostForm("usuarioID"))
	db.First(&usuario, usuarioID)

	if usuario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡Usuario no encontrado!"})
		return
	}

	//Guarda el boleto comprado
	var boleto BoletoModel
	boleto = BoletoModel{AsientoId: uint(asientoID), UsuarioId: uint(usuarioID)}
	db.Save(&boleto)

	//Actualiza la disponilidad del asiento
	db.Model(&asiento).Update("Disponible", false)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "¡Boleto comprado! ¡Gracias por su compra!"})
}