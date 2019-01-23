CREATE DATABASE IF NOT EXISTS tuticketdb;

USE tuticketdb;

CREATE TABLE evento (
	id INT NOT NULL AUTO_INCREMENT,
	nombre VARCHAR(25) NULL,
	fecha DATE NULL,
	sede_id INT NOT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY (sede_id) REFERENCES sede(id)
);

CREATE TABLE sede (
	id INT NOT NULL AUTO_INCREMENT,
	nombre VARCHAR(50) NULL,
	tipo VARCHAR(25) NULL,
	ciudad VARCHAR(25) NULL,
	pais VARCHAR(25) NULL,
	direccion VARCHAR(50) NULL,

	PRIMARY KEY(id)
);

CREATE TABLE asiento (
	id INT NOT NULL AUTO_INCREMENT,
	numero INT NULL,
	disponible BOOlEAN NULL,
	evento_id INT NULL,
	precio DECIMAL(5,2) NULL,

	PRIMARY KEY(id),
	FOREIGN KEY (evento_id) REFERENCES evento(id)
);

CREATE TABLE boleto (
	id INT NOT NULL AUTO_INCREMENT,
	asiento_id INT NULL,
	evento_id INT NULL,
	usuario_id INT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY (asiento_id) REFERENCES asiento(id),
	FOREIGN KEY (evento_id) REFERENCES evento(id),
	FOREIGN KEY (usuario_id) REFERENCES usuario(id)
);

CREATE TABLE usuario (
	id INT NOT NULL AUTO_INCREMENT,
	usuario VARCHAR(25) NULL,
	nombre VARCHAR(25) NULL,
	contrasena VARCHAR(25) NULL,

	PRIMARY KEY(id)
);