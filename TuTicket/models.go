package main

import (
	"github.com/jinzhu/gorm"
)

type (
	EventoModel struct {
		gorm.Model
		Nombre      string      `json:"nombre"`
		Descripcion string      `json:"descripcion"`
		Fecha       string      `json:"fecha"`
		SedeId      uint        `json:"sede_id"`
		Agotado     bool        `json:"agotado"`
		Asientos []AsientoModel
	}

	SedeModel struct {
		gorm.Model
		Nombre    string      `json:"nombre"`
		Tipo      string      `json:"tipo"`
		Ciudad    string      `json:"ciudad"`
		Pais      string      `json:"pais"`
		Direccion string      `json:"direccion"`
		Eventos []EventoModel
	}

	AsientoModel struct {
		gorm.Model
		Numero      uint      `json:"numero"`
		Disponible  bool      `json:"disponible"`
		EventoId    uint      `json:"evento_id"`
		LocalidadId uint      `json:"localidad_id"`
	}

	LocalidadModel struct {
		gorm.Model
		Nombre string      `json:"nombre"`
		Precio float64     `json:"precio"`
		Asientos []AsientoModel
	}

	BoletoModel struct {
		gorm.Model
		AsientoId uint      `json:"asiento_id"`
		UsuarioId uint      `json:"usuario_id"`
	}

	UsuarioModel struct {
		gorm.Model
		Usuario    string      `json:"usuario"`
		Nombre     string      `json:"nombre"`
		Contrasena string      `json:"contrasena"`
		Boletos []BoletoModel
	}
)